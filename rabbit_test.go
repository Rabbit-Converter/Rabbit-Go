package rabbit

import "testing"

func TestUni2zg(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ္ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေး ဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။", "သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝၯနေဆးၫႊန္းစာကို ဇလြန္ေဈးေဘး ဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။"},
		{"ယေဓမ္မာ ဟေတုပ္ပဘဝါ တေသံ ဟေတုံ တထာဂတော အာဟ တေသဉ္စ ယောနိရောဓေါ ဧဝံ ဝါဒီ မဟာသမဏော။", "ေယဓမၼာ ေဟတုပၸဘဝါ ေတသံ ေဟတုံ တထာဂေတာ အာဟ ေတသၪၥ ေယာနိေရာေဓါ ဧဝံ ဝါဒီ မဟာသမေဏာ။"},
	}

	for _, tc := range testCases {
		result := Uni2zg(tc.input)
		if result != tc.expected {
			t.Errorf("Uni2zg(%s) = %s; expected %s", tc.input, result, tc.expected)
		}
	}
}

func TestZg2uni(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝၯနေဆးၫႊန္းစာကို ဇလြန္ေဈးေဘး ဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။", "သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ္ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေး ဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။"},
		{"ေယဓမၼာ ေဟတုပၸဘဝါ ေတသံ ေဟတုံ တထာဂေတာ အာဟ ေတသၪၥ ေယာနိေရာေဓါ ဧဝံ ဝါဒီ မဟာသမေဏာ။", "ယေဓမ္မာ ဟေတုပ္ပဘဝါ တေသံ ဟေတုံ တထာဂတော အာဟ တေသဉ္စ ယောနိရောဓေါ ဧဝံ ဝါဒီ မဟာသမဏော။"},
	}

	for _, tc := range testCases {
		result := Zg2uni(tc.input)
		if result != tc.expected {
			t.Errorf("Zg2uni(%s) = %s; expected %s", tc.input, result, tc.expected)
		}
	}
}
