package zodiac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Zodiac struct {
	ZodiacSign     string `json:"zodiac_sign"`
	ZodiacSignInfo string `json:"zodiac_sign_info"`
	Description    string `json:"zodiac_description"`
	Compatibility  string `json:"compatibility"`
	Mood           string `json:"mood"`
	Color          string `json:"color"`
	LuckyNumber    string `json:"lucky_number"`
	LuckyTime      string `json:"lucky_time"`
}

func GetZodiac(dateBirth string) Zodiac {
	layout := "2006-01-02"
	var zodiac_sign = ""

	date, _ := time.Parse(layout, dateBirth)

	switch date.Month() {
	case time.March:
		if date.Day() < 21 {
			zodiac_sign = "Pisces"
		} else {
			zodiac_sign = "Aries"
		}
	case time.April:
		if date.Day() < 20 {
			zodiac_sign = "Aries"
		} else {
			zodiac_sign = "Taurus"
		}
	case time.May:
		if date.Day() < 21 {
			zodiac_sign = "Taurus"
		} else {
			zodiac_sign = "Gemini"
		}
	case time.June:
		if date.Day() < 21 {
			zodiac_sign = "Gemini"
		} else {
			zodiac_sign = "Cancer"
		}
	case time.July:
		if date.Day() < 23 {
			zodiac_sign = "Cancer"
		} else {
			zodiac_sign = "Leo"
		}
	case time.August:
		if date.Day() < 23 {
			zodiac_sign = "Leo"
		} else {
			zodiac_sign = "Virgo"
		}
	case time.September:
		if date.Day() < 23 {
			zodiac_sign = "Virgo"
		} else {
			zodiac_sign = "Libra"
		}
	case time.October:
		if date.Day() < 23 {
			zodiac_sign = "Libra"
		} else {
			zodiac_sign = "Scorpio"
		}
	case time.November:
		if date.Day() < 22 {
			zodiac_sign = "Scorpio"
		} else {
			zodiac_sign = "Sagittarius"
		}
	case time.December:
		if date.Day() < 22 {
			zodiac_sign = "Sagittarius"
		} else {
			zodiac_sign = "Capricorn"
		}
	case time.January:
		if date.Day() < 20 {
			zodiac_sign = "Capricorn"
		} else {
			zodiac_sign = "Aquarius"
		}
	case time.February:
		if date.Day() < 19 {
			zodiac_sign = "Aquarius"
		} else {
			zodiac_sign = "Pisces"
		}
	}

	var zodiac Zodiac
	zodiac.ZodiacSign = zodiac_sign
	zodiac.ZodiacSignInfo = GetZodiacInfo(zodiac_sign)
	mapZodiac := GetZodiacAPI(zodiac_sign)
	zodiac.Description = mapZodiac["description"]
	zodiac.Color = mapZodiac["color"]
	zodiac.Compatibility = mapZodiac["compatibility"]
	zodiac.LuckyNumber = mapZodiac["lucky_number"]
	zodiac.LuckyTime = mapZodiac["lucky_time"]
	zodiac.Mood = mapZodiac["mood"]
	return zodiac
}

func GetZodiacInfo(zodiac string) string {

	zodiacMap := map[string]string{
		"Pisces":      "Pisces is the twelfth sign of the zodiac, and it is also the final sign in the zodiacal cycle. Hence, this sign brings together many of the characteristics of the eleven signs that have come before it. Pisces, however, are happiest keeping many of these qualities under wraps. These folks are selfless, spiritual, and very focused on their inner journey. They also place great weight on what they are feeling. Yes, feelings define the Pisces zodiac sign, and it's not uncommon for them to feel their own burdens (and joys) as well as those of others. The intuition of the Pisces-born is highly evolved. Many people associate Pisces with dreams and secrets, and it's a fair association, since those born under this sign feel comfortable in an illusory world.",
		"Aries":       "Aries is the first sign of the zodiac, and that's pretty much how those born under this sign see themselves: first. Aries are the leaders of the pack, first in line to get things going. Whether or not everything gets done is another question altogether, for an Aries prefers to initiate rather than to complete. Do you have a project needing a kick-start? Call an Aries, by all means. The leadership displayed by Aries is most impressive, so don't be surprised if they can rally the troops against seemingly insurmountable odds—they have that kind of personal magnetism. An Aries sign won't shy away from new ground, either. Those born under this zodiac sign are often called the pioneers of the zodiac, and it's their fearless trek into the unknown that often wins the day. Aries is a bundle of energy and dynamism, kind of like a Pied Piper, leading people along with its charm and charisma. The dawning of a new day—and all of its possibilities—is pure bliss to an Aries.",
		"Taurus":      "Taurus, the second sign of the zodiac and the ruler of the second house, is all about reward. Unlike the Aries love of the game, the typical Taurus personality loves the rewards of the game. Think physical pleasures and material goods, for those born under this sign revel in delicious excess. This zodiac sign is also tactile, enjoying a tender, even sensual, touch. Taurus zodiac sign adores comfort and likes being surrounded by pleasing, soothing things. Along these lines, they also favor a good meal and a fine wine. The good life in all its guises, whether it's the arts or art of their own making (yes, these folks are artistic as well), is heaven on Earth to the Taurus-born.",
		"Gemini":      "Gemini is the third sign of the zodiac, and those born under this sign will be quick to tell you all about it. That's because they love to talk! It's not just idle chatter with these folks, either. The driving force behind a Gemini zodiac sign's conversation is their mind. Ruling the third house, the Gemini-born are intellectually inclined, forever probing people and places in search of information.The more information a Gemini collects, the better. Sharing that information later on with those they love is also a lot of fun, for Geminis are supremely interested in developing their relationships. Dalliances with those of this astrology sign are always enjoyable, since Geminis are bright, quick-witted, and the proverbial life of the party. Even though their intellectual minds can rationalize forever and a day, Geminis also have a surplus of imagination waiting to be tapped. Can a Gemini be boring? Never!",
		"Cancer":      "Cancer, the fourth sign of the zodiac, is all about home. Those born under this horoscope sign are ‘roots' kinds of people, and take great pleasure in the comforts of home and family.Cancers are maternal, domestic, and love to nurture others. More than likely, their family will be large, too—the more, the merrier! Cancers will certainly be merry if their home life is serene and harmonious. Traditions are upheld with great zest in a Cancer's household, since, as the rulers of the fourth house of home and memory, this zodiac sign prizes family history and loves communal activities. They also tend to be patriotic, waving the flag whenever possible. A Cancer's good memory is the basis for stories told around the dinner table, and don't be surprised if these folks get emotional about things. Those born under this sign wear their heart on their sleeve, which is just fine by them.",
		"Leo":         "Leo is the fifth sign of the zodiac. These folks are impossible to miss since they love being center stage. Making an impression is Job #1 for Leos, and when you consider their personal magnetism, you see the job is quite easy. Leos are an ambitious lot, and their strength of purpose allows them to accomplish a great deal. The fact that this horoscope sign is also creative makes their endeavors fun for them and everyone else.It's quite common to see a Leo on stage or in Hollywood since these folks never shy away from the limelight. They are also supremely talented and have a flair for the dramatic. Warmth and enthusiasm seem to seep from every Leo pore, making these folks a pleasure to be around. They do love pleasure and being the center of attention!",
		"Virgo":       "Virgo is the sixth sign of the zodiac, to be exact, and that's the way Virgos like it: exacting. Those born under this horoscope sign are forever the butt of jokes for being so picky and critical (and they can be), but their ‘attention to detail' is for a reason: to help others. Virgos, more than any other zodiac sign, were born to serve, and it gives them great joy. They are also tailor-made for the job, since common Virgo traits are being industrious, methodical, and efficient. The sense of duty borne by these folks is considerable, and it ensures that they will always work for the greater good.",
		"Libra":       "Libra is the seventh sign of the zodiac, and it's at this point in the zodiac that we start to see a shift. While the first six signs of the zodiac focus on the individual, the last six focus on the individual's contact with others and with the world. The Libra zodiac sign is first and foremost focused on others and how they relate to them. We can call this the sign of Partnership with a capital 'P' because these folks do not want to be alone! For a Libra, everything is better if it's done as a pair. Libras are good when paired up, too, since they epitomize balance, harmony, and a sense of fair play. While they are true team players at work, their favorite partnership is at home: marriage. Libras feel most complete when they are coupled up with their lover, forever.",
		"Scorpio":     "Scorpio is the eighth sign of the zodiac, and that shouldn't be taken lightly—nor should Scorpios! Those born under this sign are dead serious in their mission to learn about others. There's no fluff or chatter for Scorpios, either; these folks will zero-in on the essential questions, gleaning the secrets that lie within. The Scorpio zodiac sign concerns itself with beginnings and endings, and is unafraid of either. They also travel in a world that is black and white and has little use for gray. The curiosity of Scorpios is immeasurable, which may be why they are such adept investigators.The folks with a Scorpio horoscope sign love to probe and know how to get to the bottom of things. The fact that they have a keen sense of intuition certainly helps.",
		"Sagittarius": "Sagittarius, the ninth sign of the zodiac, is the home of the wanderers of the zodiac. It's not a mindless ramble for these folks, either. Sags are truth-seekers, and the best way for them to do this is to hit the road, talk to others and get some answers.Knowledge is key to these folks since it fuels their broad-minded approach to life. Those born with a Sagittarius zodiac sign are keenly interested in philosophy and religion, and they find that these disciplines aid their internal quest. At the end of the day, what Sagittarius wants most is to know the meaning of life, and to accomplish this while feeling free and easy.",
		"Capricorn":   "Capricorn, the tenth sign and mountain goat of the zodiac, is all about hard work. Those born under this sign are more than happy to put in a full day at the office, realizing that it will likely take a lot of those days to get to the top. That's no problem, since Capricorns are both ambitious and determined: they will get there. Life is one big project for these folks, and they adapt to this by adopting a businesslike approach to most everything they do. Capricorns are practical as well, taking things one step at a time and being as realistic and pragmatic as possible. Those with a Capricorn zodiac sign are extremely dedicated to their goals, almost to the point of stubbornness. Those victories sure smell sweet, though, and that thought alone will keep Capricorns going.",
		"Aquarius":    "Aquarius is the eleventh sign of the zodiac, and Aquarians are the perfect representatives for the Age of Aquarius. Those born under this horoscope sign have the social conscience needed to carry us into the new millennium. Those of the Aquarius zodiac sign are humanitarian, philanthropic, and keenly interested in making the world a better place. Along those lines, they'd like to make the world work better, which is why they focus much of their energy on our social institutions and how they work (or don't work). Aquarians are visionaries, progressive souls who love to spend time thinking about how things can be better. They are also quick to engage others in this process, which is why they have so many friends and acquaintances. Making the world a better place is a collaborative effort for Aquarians.",
	}

	return zodiacMap[zodiac]
}

//API
func GetZodiacAPI(zodiac string) map[string]string {
	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=" + zodiac + "&day=today"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Host", "sameer-kumar-aztro-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "50be42a795mshb68f9a3788bc691p190dffjsnd755284183c7")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	m := make(map[string]string)
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return m
}
