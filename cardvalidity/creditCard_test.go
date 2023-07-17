package cardvalidity

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
	"testing"
)

func TestCreditCardValidity(t *testing.T) {

	testCases := `1182387522195848,Invalid
4898285859544556,Valid
3533946521218854,Invalid
2178579985193175,Invalid
9691928949545344,Invalid
1327576477684519,Invalid
6885867822596993,Valid
1958129523778579,Invalid
2357676157819124,Invalid
9832746248713726,Invalid
3673159777236652,Invalid
8626186974574971,Invalid
9687622296992497,Invalid
6731749895254584,Valid
8231635922318254,Invalid
7728878259735616,Invalid
3347275338764373,Invalid
6624557432269847,Valid
3164653818478977,Invalid
7683817293887423,Invalid
4654491168789767,Valid
6942469919536219,Valid
8434524674271379,Invalid
6619249165433473,Valid
8842787232483367,Invalid
5349898497837349,Valid
4841565975496529,Valid
7635659522159832,Invalid
6699889899699968,Valid
5274676861386577,Valid
7261479482325831,Invalid
9855821811363989,Invalid
5462941623272486,Valid
2168457338741692,Invalid
3493828267535654,Invalid
7688277563695358,Invalid
4621162653647299,Valid
5588937472734175,Valid
6313634439334582,Valid
2621731928824298,Invalid
9356326214767474,Invalid
6399396262113367,Valid
7326622854597675,Invalid
2179646384144599,Invalid
8723731421194264,Invalid
9893925198222769,Invalid
8493394862176119,Invalid
9182146786584817,Invalid
7865278423923993,Invalid
5437343432579992,Valid
61234-567-8912-3456,Invalid
6123-4567-8912-3456,Valid
4123356789123456,Valid
5133-3367-8912-3456,Invalid
5123-3567-8912-3456,Valid
4443-3366-6777-3333,Invalid
4443-3366-6777-3331,Valid
5123 - 3567 - 8912 - 3456,Invalid`
	r := csv.NewReader(strings.NewReader(testCases))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		got := ValidateCard(strings.Trim(record[0], " "))
		if got != strings.Trim(record[1], " ") {
			t.Errorf("For card : %v expected: %v and got: %v", record[0], record[1], got)
		}
	}

}
