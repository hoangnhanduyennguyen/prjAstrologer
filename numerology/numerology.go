package numerology

import (
	"strconv"
	"strings"
)

type Numerology struct {
	LifePathNumber  int    `json:"life_path_number"`
	LifePathMeaning string `json:"life_path_meaning"`
	DestinyNumber   int    `json:"destiny_number"`
	DestinyMeaning  string `json:"destiny_meaning"`
}

func GetNumerology(dobInput, firstName, lastName string) Numerology {
	var numerology Numerology
	lifePath := GetLifePathNumber(dobInput)
	numerology.LifePathNumber = lifePath
	numerology.LifePathMeaning = GetLifePathMeaning(lifePath)
	destinyNumber := GetDestinyNumber(firstName, lastName)
	numerology.DestinyNumber = destinyNumber
	numerology.DestinyMeaning = GetDestinyNumberMeaning(destinyNumber)

	return numerology
}

func GetLifePathNumber(dobInput string) int {
	var lifePath int

	dob := []rune(dobInput)
	var dobRune []rune = dob
	for {
		for _, val := range dobRune {
			num, _ := strconv.Atoi((string(val)))
			lifePath += num

		}
		if lifePath <= 9 || lifePath == 11 || lifePath == 22 || lifePath == 33 {
			break
		}
		s := strconv.Itoa(lifePath)
		dobRune = []rune(s)
		lifePath = 0
	}

	return lifePath
}

func GetLifePathMeaning(num int) string {
	lifePathMap := map[int]string{
		1:  "The life path number 1 is ruled by the Sun. They are a born leader, ambitious for goal and hard workers. They are very protective of the family and love. Although, they want respect and attention from the family and others. They should control their anger, pride, zealousness, and focus on the goals, which they want to achieve. Sometimes, they are very critical of others and become irritated, when things do not happen according their desires. As They have strong desires to be number one or aggressive in the life. They should avoid laziness, egotistical nature and do their best efforts, because no one can do so best as they can!",
		2:  "The life path number two is influenced by the Moon. So, they are polite, peacemaker, shy, sensitive and gentle person, who like to help others. They are lovers of music, arts, harmonious environment and very loyal to the lover. Their sensitivity and feelings make them sad or happy sometimes, but these qualities give them a power of healing also! Other people want their company due to the peacemaker behavior and a good sense of humor. They are the perfectionist and have their own silent ways to complete the tasks, but often they do not get credit for the works, which they had done. They have a quality to change the life of others and themselves!",
		3:  "The life path number 3 is impacted by the Jupiter. They are moody person, who often ignore responsibilities and believe to go easy with any situation. They have a great quality of creativity and self-expression, and most of the writers, artists, actors, singers, designer, brokers and publishers are born under the life path number three. They are positive all about the life, have a deep desire for the luxurious life, and of course, they spend the life with the luxury and harmony. As they are so positive, that they do not plan their future seriously, therefore, not so lucky in the monetary matters. They are social people and may get huge fame for their work.",
		4:  "The life path number 4 is ruled by the Uranus or Raahu. They are persons, who take responsibilities very seriously for their family or works. They are the good planner for their life, business and future, because they are so practical, hard working and down to earth person. They are courageous and committed person, and try to establish the foundation for their business and life. They are honest and accept honesty from others. They are loyal and love full of the life partner and children. The fours like to be a lawyer, doctor, engineer, politician, manager, electrician or accountant because of discipline and hardworking nature.",
		5:  "The life path number five is influenced by the Mercury. They are persons, who ever seek the changes and freedom in the work and life. They are persuasive to others, restless for work and the lover of adventures. They can be in any field, which requires communication skills, like salesmanship, publishing, advertising or promotional activities. They always take risks and do work with the discipline and step by step, but hate the routine and repetition also. They are multi-talented and have many abilities. Because of the frequent change of the nature, they leave the goals in the middle path and seek new adventures in the works, so many of the tasks, remain unfinished.",
		6:  "The life path number 6 is ruled by the Venus. They are very kind, loving, compassionate, responsible and loyal to the family, friends and others. They like to help others at any cost, due to this nature, they always ready to sacrifice, but this nature makes them an interrupter in the life of others. They are family and home-oriented, who love to be honest and aware full of them. The person of number six likes to be a teacher, broker, caretaker, hotelier or a club person. They are humble, generous and attractive, but sometimes have ego inside, so they should not try to deny that who are they and should try to work for themselves and then seek to help others.",
		7:  "The life path number seven is influenced by the Ketu or Neptune. They are persons, who search the truth about spirituality and other people. They are like to be alone, even in the crowd because of lack of faith in others. They are the seeker of knowledge, analytical thinkers, reserved, deep thinkers and have a charming personality. They are not so socially and flexible in the society, because of an introvert nature. The number seventh persons are holy and spiritual, but due to introvert nature, they cannot manage balance in marriage and social life. Sometimes, they feel jealous of others, thus selfishness or ego drives them on a path, which others do not like.",
		8:  "The life path number 8 is impacted by the Saturn. They are persons, who have great abilities of leadership, political skills and decisive and commanding. They can be the good businessmen, politicians, judges and builders. They are God gifted to earn money and fame in the world, but they do not like the advice of others. They should control their attitude, because sometimes, they become stubborn, intolerant and impatient. They are very workaholic and it may cause the health problems. They like a status in the society and life, therefore they believe their ways are right, which make family, friends, and others uncomfortable and conflicted.",
		9:  "The life path number 9 is ruled by the Mars. They are the leader, friendly and angry person and like to help others with the time and money, but often feel ignored or unloved by the parents. Sometimes, they feel dissatisfaction with the life and the results of the works. Often, they get money and fame in the unexpected ways, but they do not manage their money or fame seriously. They are imaginative and creative and have artistic, healing and writing talents, and likely to be a fashion designer, interior decorator or a writer. They are romantic, but moody or ungrateful too. They give and sacrifice for others, but when they unsuccessful, they try to blame others.",
		11: "The life path number 11 is influenced by the Moon. They are spiritual and very intuitive person, who are sensitive and gentle too. They like to help others. The 11 is a master builder number. The 11/2 is an inspirational teacher. Here, it can be assumed the 2 is a single digit life path number of 11. They are deep thinkers, spiritually aware, dreamer and loyal to others. Sometimes, they may face fears or phobia, and become moody and nervous. They are faithful to life partner or business partner, and have blind faith in others too. Other people want their company due to the peacemaking behavior and the good sense of humor. They are the perfectionist and have own silent ways to complete the tasks, but often they do not get credit for the works they have done previously.",
		22: "The life path number 22 is ruled by the Uranus or Rahu. The 22 is a master builder number. The 22/4 is called a master healer. Here, it can be assumed the 4 is a single digit life path number of 22. They are very practical in the life, who take own responsibility seriously for the family and work. They are born under most powerful and successful life path number, so they can get success with their plans and hardworking nature, but they always have a fear of failure. They are honest people and accept honesty in others also. They are loyal and love full of life partner and children. They like to be in politics, law or business. They are romantic, but moody or ungrateful too. Their thoughts have great powers, so be careful with own thoughts!",
		33: "The life path number 33 is ruled by the Venus. The 33 is a master builder number. The 33/6 is a cosmic parent. Here, it can be assumed the 6 is a single digit life path number of 33. They are very spiritual and down to earth person. They are the spiritual leaders and teachers, who can make the world peace. They can get fame with their kind work and compassion in the society. They are very loving, responsible and loyal to the family, friends and others. They like to help others anytime, and for this they always ready to make sacrifices also. The person of number 33 likes to be a teacher, broker, holy person, caretaker, hotel or club person. They are humble, generous and attractive, but sometimes have the ego inside.",
	}
	return lifePathMap[num]
}

func GetDestinyNumber(first string, last string) int {
	var destinyNumber int
	var fullname = []rune(strings.TrimSpace(strings.ToLower(first)) + " " + strings.TrimSpace(strings.ToLower(last)))

	for i := 0; i < len(fullname); i++ {
		if fullname[i] >= 115 {
			fullname[i] = fullname[i] - 66
		} else {
			if fullname[i] >= 106 {
				fullname[i] = fullname[i] - 57
			} else {
				if fullname[i] >= 61 {
					fullname[i] = fullname[i] - 48
				}
			}
		}

	}
	for {
		for _, val := range fullname {
			num, _ := strconv.Atoi((string(val)))
			destinyNumber += num

		}
		if destinyNumber <= 9 || destinyNumber == 11 || destinyNumber == 22 || destinyNumber == 33 {
			break
		}
		s := strconv.Itoa(destinyNumber)
		fullname = []rune(s)
		destinyNumber = 0
	}

	return destinyNumber
}

func GetDestinyNumberMeaning(num int) string {
	destinyNumberMap := map[int]string{
		1:  "Destiny number 1 is influenced by the qualities of the planet Sun, like firm determination, bravery and leadership qualities. The no. 1 people are original thinker and originator of all actions. They have a quality to control everything around them. Sometimes, they become stubborn and angry, when things do not happen, according to their desire. They are a real and original thinker, creative and dominate persons. They can play the good role as the inventor, leader, explorer and a head of the family. The number one represents the card Magician.",
		2:  "The destiny number 2 people are ruled and influenced by the planet Moon. The number two represents sensitivity, imagination and dream oriented nature. They are such a romantic and cooperative person, other people want to live with them, people know them as a real peacemaker. They are diplomat persons with the duo personality. The person of no. 1 is fond of music and may become a better singer. They love to travel and home too. Many times, they have unknown fears in the life, so they adopt a secret behavior, thus people cannot judge their next move. The more energy of number two may make them shy, oversensitive and fantasy dreamer. The number two represents the card High Priestess.",
		3:  "The expression number 3 has qualities of the planet Jupiter, and influenced by the Jupiter during the whole life. They are social persons, who are creative, communicative and dramatic. The number 3 represents artistic talents, charismatic personality and cheerful behavior. They are the religious, truthful, highly educated or highly skilled person. The no. 3 can become the great ideal for others, and love to travel, learn the new ways of joy and happiness. They trust about total freedom in every the aspect of life, especially the freedom of speech. The person of destiny number three can become the lawyer, artist, writer or publisher. The number three represents the card Empress.",
		4:  "The Uranus or Rahu is known for sudden and unexpected events in the mysterious ways. People of number three are influenced by the Uranus or Rahu in their life. They may face many unexpected events in the life, due to the planet. The fourth is a practical person, who believes in individualism, tolerance, and originality. The secretive nature and unexpected behavior to others are the main traits of number four. They always make plans for the foundation for the future with the weirdest ideas. They are called systematic and trustworthy. They have a long list of the likes and dislikes, they should avoid stubborn and rigid behavior to others. The number four represents the card Emperor.",
		5:  "People of expression number five have traits of the planet Mercury, such as fast movement, versatile nature, communication skill and multi-talents. These qualities give them the opportunities to become a salesperson, actor, writer, media and commercial post related person and sales or communication expert. The person who has the number 5 as destiny number is a clever, so has the quicker response behavior and great ideas for any problem. They love freedom, adventure, and work in the higher risk. Some people of number five are highly workaholic and restless, due to this, they may suffer from nerve related problems. Due to analytical nature, they disappoint the persons, who nearby them. The number five represents the card Hierophant.",
		6:  "The persons who have destiny no. Six, are influenced by the Venus, a planet of love, romance, beauty, art and truth. So, they are loving persons with a romantic image. They strongly trust in the truth, justice, and humanity. They are born teacher or healer, who always ready to help others with their counseling skills. The person of number six is the divine lover and always stands for family and friends. They love luxury and harmony in the life. Number six person may become good singer, counselor, teacher or art related expert. The number six represents the card Lovers.",
		7:  "The number seven persons remain under the influence of Neptune or Ketu planet. This planet represents spirituality, philosophy and mystery. They are always searching the answers of questions, which they face during the life journey. They are forced by number seven to make illusion and delusion. They have mysterious nature, they do dream a lot than others. And, always investigate the reasons of the happenings. They often feel, that they are gifted with some type of clairvoyance and intuition powers. They are introvert person and often do not share feelings and problems with others. The number seven represents the card Chariot.",
		8:  "The expression number Eight's persons are under the influence of the planet Saturn. This planet represents stability, true judgement and great responsibility. They have born qualities to manage financial things and politics, indeed, the person with number eight may become a great businessman, a leader or a game changer in the finance or politics. They have a good self-control and determination power, with the reserved nature. They can make balance everything in the life and between situations. The number eight represents the card Justice.",
		9:  "People who have destiny number 9, are influenced by the qualities of the planet Mars, like they are aggressive persons, who do not fear to take risks, they are the courageous persons with traits of humanity and kindness. They are always ready to help others, when needed. They get all things, what is determined by taking any levels of the risks. They trust everyone and often defeated by the hidden enemies, although they cannot be defeated, if they know enemies. The anger and impatience are the worst part of their nature. The number nine represents the card Hermit.",
		11: "This number 11, which is a master builder number, and 2 is standing for a single digit number are ruled by the planet Moon. Because, the 11 number called master number, it means, it has double and special powers other than 2. It represents the high sensitivity, spiritual and high imagination powers. The 11 number persons may become a intuitive and psychic person. Because of the Moon, they are very soft to others, and have blind trust in others. And, in a result, often get treachery and hidden dangers from others. They have a very secretive nature, people cannot understand them. The number eleven represents the card Lion Muzzled.",
		22: "The 22 is called master number, the 4 is a single number, which is influenced by the Uranus/Rahu. Due to the planet traits, they often live in the dreams. They have own world of illusion, often awake after falling into the danger. The unexpected events are the main traits of number 22 during the whole life. They may face misfortune, hidden enemies and hidden dangers and always feel surrounded. Although, the master number gives more powers, like they are very trustworthy, good manager, energetic, capable and systematic than number two person. They always make plans for the better future, but they should open eyes about the warning of future events, because of the illusion nature.",
		33: "The 33 is known as a master builder number, and the 6 is a single digit number, which is ruled and influenced by the planet Venus. Due to double power of 33 master number, this is a fortunate or lucky number. This brings fortune in the love matters, finance and money. They are the true and charming lover of the partner. They are overconfident, but skilled and highly creative person. The amazing teachers, counselors, and healers were born under the number 33.",
	}
	return destinyNumberMap[num]
}
