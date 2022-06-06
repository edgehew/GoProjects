package main 

import (
	"math/rand"
	"time"
)

var (
	insultList = []string {
		"I was today years old when I realized I didn’t like you.",
		"Someday you’ll go far. And I really hope you stay there.",
		"Oops, my bad. I could’ve sworn I was dealing with an adult.",
		"I love what you’ve done with your hair. How do you get it to come out of your nostrils like that?",
		"Remember that time you were saying that thing I didn’t care about? Yeah, that is now.",
		"You’re the reason God created the middle finger.",
		"I’m busy right now, can I ignore you another time?",
		"Oh, you don’t like being treated the way you treat me? That must suck.",
		"I wish I had a flip phone, so I could slam it shut on this conversation.",
 		"N’Sync said it best, 'BYE, BYE, BYE!'",
		"I’ve been called worse things by better men.",
		"You’re a gray sprinkle on a rainbow cupcake.",
		"Your secrets are always safe with me. I never even listen when you tell me them.",
		"You bring everyone so much joy! You know, when you leave the room. But, still.",
		"How many licks until I get to the interesting part of this conversation?",
		"Keep rolling your eyes, you might eventually find a brain.",
		"Your face makes onions cry.",
		"Did I invite you to the barbecue? Then why are you all up in my grill?",
		"Our kid must have gotten his brain from you! I still have mine.",
		"You have so many gaps in your teeth it looks like your tongue is in jail.",
		"If your brain was dynamite, there wouldn’t be enough to blow your hat off.",
		"You are more disappointing than an unsalted pretzel.",
		"It’s impossible to underestimate you.",
		"Wow, your maker really didn’t waste time giving you a personality, huh?",
		"Her teeth were so bad she could eat an apple through a fence.",
		"I’ll never forget the first time we met. But I’ll keep trying.",
		"Oh, I’m sorry. Did the middle of my sentence interrupt the beginning of yours?",
		"Hold still. I’m trying to imagine you with personality.",
		"I’m not insulting you, I’m describing you.",
		"You are the human version of period cramps.",
		"You lack brains so much that you can float on water.",
		"I will insult you. But then I’ll have to explain later. It’s okay.",
		"I don’t have time or crayons to explain to you.",
		"We do not accept crayon-written resumes."
		"You know nothing. In fact, you know less than nothing. ‘Cause if you know you know nothing That will be something.",
		"I expected an intellectual conversation. But it seems that no one is there.",
	}
)

func getRandomInsult() (s string) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return insultList[rand.Intn(len(insultList))]
}