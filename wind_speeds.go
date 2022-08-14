package bbcweather

type WindType struct {
	BeaufortNumber  int
	MaxWindSpeedMph int
	Summary         string
	Description     string
}

var winds = []WindType{
	{BeaufortNumber: 0, MaxWindSpeedMph: 1, Summary: "Calm", Description: "Calm wind. Smoke rises vertically with little if any drift."},
	{BeaufortNumber: 1, MaxWindSpeedMph: 3, Summary: "Light Air", Description: "Direction of wind shown by smoke drift, not by wind vanes. Little if any movement with flags. Wind barely moves tree leaves."},
	{BeaufortNumber: 2, MaxWindSpeedMph: 7, Summary: "Light Breeze", Description: "Wind felt on face. Leaves rustle and small twigs move. Ordinary wind vanes move."},
	{BeaufortNumber: 3, MaxWindSpeedMph: 12, Summary: "Gentle Breeze", Description: "Leaves and small twigs in constant motion. Wind blows up dry leaves from the ground. Flags are extended out."},
	{BeaufortNumber: 4, MaxWindSpeedMph: 18, Summary: "Moderate Breeze", Description: "Wind moves small branches. Wind raises dust and loose paper from the ground and drives them along."},
	{BeaufortNumber: 5, MaxWindSpeedMph: 24, Summary: "Fresh Breeze", Description: "Large branches and small trees in leaf begin to sway. Crested wavelets form on inland lakes and large rivers."},
	{BeaufortNumber: 6, MaxWindSpeedMph: 31, Summary: "Strong Breeze", Description: "Large branches in continuous motion. Whistling sounds heard in overhead or nearby power and telephone lines. Umbrellas used with difficulty."},
	{BeaufortNumber: 7, MaxWindSpeedMph: 38, Summary: "Near Gale", Description: "Whole trees in motion. Inconvenience felt when walking against the wind."},
	{BeaufortNumber: 8, MaxWindSpeedMph: 46, Summary: "Gale", Description: "Wind breaks twigs and small branches. Wind generally impedes walking."},
	{BeaufortNumber: 9, MaxWindSpeedMph: 54, Summary: "Strong Gale", Description: "Structural damage occurs, such as chimney covers, roofing tiles blown off, and television antennas damaged. Ground is littered with many small twigs and broken branches."},
	{BeaufortNumber: 10, MaxWindSpeedMph: 63, Summary: "Whole Gale", Description: "Considerable structural damage occurs, especially on roofs. Small trees may be blown over and uprooted."},
	{BeaufortNumber: 11, MaxWindSpeedMph: 75, Summary: "Storm Force", Description: "Widespread damage occurs. Larger trees blown over and uprooted."},
	{BeaufortNumber: 12, MaxWindSpeedMph: -1, Summary: "Hurricane Force", Description: "Severe and extensive damage. Roofs can be peeled off. Windows broken. Trees uprooted. RVs and small mobile homes overturned. Moving automobiles can be pushed off the roadways."},
}

func GetWindTypeFromSpeed(speedMph int) WindType {
	for _, wind := range winds {
		if speedMph <= wind.MaxWindSpeedMph {
			return wind
		}
	}
	return winds[len(winds)-1]
}
