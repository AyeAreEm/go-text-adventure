package main

import (
	"fmt"
	"os"
	"time"
)

var money int = 500
var silver string = "common"
var steel string  = "common" 
var health int = 100

func main() {
	fmt.Println("You are in a forest. To the north is a castle. To the south is a village. Where do you go?")
	var start string
	fmt.Scanln(&start)

	if start == "north" {
		begNorth()

	} else if start == "south" {
		begSouth()
	} else {
		fmt.Println("That is not an option.")
	}
}

func begNorth() {
	fmt.Println("You head north. There is a massive gate and a guard guarding it. Do you talk to the guard or go to the village. (guard / village)")
	var gOrV string
	fmt.Scanln(&gOrV)

	if gOrV == "guard" {
		fmt.Println("Guard: What do you want Witcher. We don't want your kind here.")
		time.Sleep(2 * time.Second)
		fmt.Println("You: May I go inside the castle? It's important.")
		time.Sleep(2 * time.Second)
		fmt.Println("Guard: Oh piss off Witcher! As I said, you are not welcome here.")
		time.Sleep(2 * time.Second)
		fmt.Println("What do you do? $100 crowns or use Axii (mind control)? (crowns / axii)")
		var hOrA string
		fmt.Scanln(&hOrA)

		if hOrA == "crowns" {
			money = money - 100
			fmt.Println("You give 100 crowns to the guard. You now have", money)
			time.Sleep(2 * time.Second)
			fmt.Println("Guard: What do you take of me Witcher? A fucking fool? ")
			time.Sleep(2 * time.Second)
			fmt.Println("You leave.")
			time.Sleep(2 * time.Second)
			begSouth()
		} else if hOrA == "axii" {
			fmt.Println("You: Open the gate")
			fmt.Println("Guard: Yes Witcher. Right away.")
			castle()
		}

	} else if gOrV == "village" {
		begSouth()
	}
}

func begSouth() {
	fmt.Println("You walk into the village. You talk to a peasant. What do you say? Ask about the castle or go to the castle by yourself? (ask / castle)")
	var aOrC string
	fmt.Scanln(&aOrC)

	if aOrC == "ask" {
		fmt.Println("You: How do I get into the castle? ")
		time.Sleep(2 * time.Second)
		fmt.Println("Peasant: What's in it for me Witcher?")
		time.Sleep(2 * time.Second)
		fmt.Println("What do you offer the peasant? $50 crowns or ask for a contract? Or use Axii (mind control)? (crowns / contract / axii)")
		var fCOrA string
		fmt.Scanln(&fCOrA)

		if fCOrA == "crowns" {
			money = money - 50
			fmt.Println("You: Hmmph. Fine, how about 50 crowns?")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: That'll do witcher.")
			time.Sleep(2 * time.Second)
			fmt.Println("You gave the peasant 50 crowns and now you have", money)
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: Awhile ago, the King's child disappeared for around a fortnight and was found dead in the river north-east of here. Since he was last seen in the castle before disappearing...")
			time.Sleep(2 * time.Second)
			fmt.Println("You: Hm. So there should be something near the river to get me inside the castle.")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: Precisely!")
			enterCastle()

		} else if fCOrA == "contract" {
			fmt.Println("You: How about I do a contract for you? Any monsters that need to be taken care of?")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: I do in fact have a contract for you. I have heard word of a bird like creature terrorizing farmers in the village. Killing cows, famers and even soldiers. Slay it and I will help you.")
			time.Sleep(2 * time.Second)
			fmt.Println("You: I will kill this beast for you but I still require coin for my services. Witcher's code if you will.")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: Fine Witcher! I will give you 100 crowns if you slay it.")
			time.Sleep(2 * time.Second)
			fmt.Println("You: Alright. Where is its nest?")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: It's west of here.")
			griffin()

		} else if fCOrA == "axii" {
			fmt.Println("You: Tell me how to get into the castle.")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: Yes, of course Witcher.")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: Awhile ago, the King's child disappeared for around a fortnight and was found dead in the river north-east of here. Since he was last seen in the castle before disappearing...")
			time.Sleep(2 * time.Second)
			fmt.Println("You: Hm. So there should be something near the river to get me inside the castle.")
			time.Sleep(2 * time.Second)
			fmt.Println("Peasant: Precisely!")
			enterCastle()
		}

	} else if aOrC == "castle" {
		begNorth()
	}
}

func griffin() {
	fmt.Println("You start riding west on Roach.")
	time.Sleep(2 * time.Second)
	fmt.Println("You see the nest and a griffin sleeping. What do you do? Stab the neck or use Aard (telekentic blast)? (stab / aard)")
	var sOrB string
	fmt.Scanln(&sOrB)

	if sOrB == "stab" {
		fmt.Println("You quietly sneak up to the griffin and stab it right in the neck. It squeals in pain. You push harder into the griffin, destroying it's voice box. The griffin is trying to scream in pain but literally can't and just makes gusts of air as it slowly dies.")
		time.Sleep(2 * time.Second)

	} else if sOrB == "aard" {
		fmt.Println("You cast Aard. It awakens the griffin but it has troubles getting up. You lunge at it with your", silver, " silver sword. The griffin screams in pain and knocks you back using its wings. You cast Aard again, knocking the griffin back as well. With the griffin lying on its back, you take your sword and stab it right in the stomach, and dragged your sword from the top to bottom... killing the griffin.")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("You head back to village.")
	time.Sleep(2 * time.Second)
	fmt.Println("You: I killed your fucking griffin. Give me the coin and tell me how to get into the castle!")
	time.Sleep(2 * time.Second)
	fmt.Println("Peasant: Fair is fair Witcher. Here is the 100 crowns and for how to get into the castle....")
	fmt.Println("You: Mhm?")
	time.Sleep(2 * time.Second)
	fmt.Println("Peasant: Awhile ago, the King's child disappeared for around a fortnight and was found dead in the river north-east of here. Since he was last seen in the castle before disappearing...")
	time.Sleep(2 * time.Second)
	fmt.Println("You: Hm. So there should be something near the river to get me inside the castle.")
	time.Sleep(2 * time.Second)
	fmt.Println("Peasant: Precisely!")
	money = money + 100
	enterCastle()
}

func enterCastle() {
	right := func()  {
		fmt.Println("You go to the right. You see a monster. You examine the monster quickly and realize that it's a wraith.")
		time.Sleep(2 * time.Second)
		fmt.Println("What do you do? Charge at the wraith or use Yrden (a magic trap)? (charge / yrden)")
		var cOrY string
		fmt.Scanln(&cOrY)

		if cOrY == "charge" {
			fmt.Println("You charge at the wraith with your", silver, " silver sword.")
			time.Sleep(2 * time.Second)

			if silver == "rare" {
				fmt.Println("You charge and you slash the wraith. The wraith teleports behind you and knocks you down onto the floor. It grabs your legs and flings you to the wall. The wraith starts teleporting right in front of you... you lift your sword onto your chest with the blade pointing away from you. The wraith teleported right into your sword, killing it.")
				time.Sleep(2 * time.Second)
				health = health - 70
				fmt.Println("You are nearly dead but not there yet. You're health is", health)
				time.Sleep(2 * time.Second)

			} else if silver == "common" {
				fmt.Println("You charge and barely graze the wraith. The wraith teleports behind you and knocks you down onto the floor. It grabs your legs and flings you to the wall. It then teleports right in front of you and... kills you.")
				os.Exit(0)
			}

		} else if cOrY == "yrden" {
			fmt.Println("You cast yrden. The wraith teleports behind you, trying to attack but you trapped it inside. You slash left... slash right... and stab right in the middle, killing the wraith.")
			time.Sleep(2 * time.Second)
		}

		fmt.Println("You see a rope leading up into a well. You climb it and enter the castle.")
		castle()

	}

	fmt.Println("You call Roach and start riding to the river.")
	time.Sleep(2 * time.Second)
	fmt.Println("You see a cave like entrance underwater. You proceed to enter.")
	time.Sleep(2 * time.Second)
	fmt.Println("The cave splits off into two tunnels. The one on the left has a light at the end. The right one leads into darkness. Which one do you go through? (left / right)")
	var lOrR string
	fmt.Scanln(&lOrR)

	if lOrR == "left" {
		fmt.Println("You go to the left. You realize that there is nothing here and it just lead you back outside the cave. But you do find a rare tier silver sword. Do you take it? (yes / no)")
		var yOrN string
		fmt.Scanln(&yOrN)

		if yOrN == "yes" {
			silver = "rare"
			fmt.Println("You take the steel sword and head back from where you came.")
			right()
		} else if yOrN == "no" {
			fmt.Println("You don't take the steel sword. You start heading back from where you came.")
			right()
		}
	}

	if lOrR == "right" {
		right()
	}

}

func castle() {
	fmt.Println("As you enter the castle, you see the King talking to his subjects. What do you do? Wait for the discussion to finish or push the subjects away and talk to the King? (wait / push)")
	var wOrP string
	fmt.Scanln(&wOrP)

	if wOrP == "wait" {
		fmt.Println("You wait for the discussion to finish. The King notices you and quickly finishes the discussion.")
		time.Sleep(2 * time.Second)
		discussion()
	} else if wOrP == "push" {
		fmt.Println("You push aside the subjects. They look at you with disbelief. You reply to their disbelief with a death stare.")
		time.Sleep(2 * time.Second)
		discussion()
	}
}

func discussion() {
	fmt.Println("King: Well hello Witcher. What are you here for?")
	time.Sleep(2 * time.Second)
	fmt.Println("You: You know what I'm here for!")
	time.Sleep(2 * time.Second)
	fmt.Println("King: I have no clue Witcher. Please tell me.")
	time.Sleep(2 * time.Second)
	fmt.Println("You: Where is Ciri?!")
	time.Sleep(2 * time.Second)
	fmt.Println("King: I still don't know what you are talking about Witcher!")
	time.Sleep(2 * time.Second)
	fmt.Println("What do you do to make the King talk? Use Axii, ask for contract or threaten? (axii / contract / threaten)")
	var aCOrT string
	fmt.Scanln(&aCOrT)

	if aCOrT == "axii" {
		fmt.Println("You: Tell me where Ciri is.")
		time.Sleep(2 * time.Second)
		fmt.Println("King: Yes Witcher. She is locked up in my house. She is healthy and fine. Just alone in inside.")
		time.Sleep(2 * time.Second)
		fmt.Println("Some of the guards realize that you casted Axii on the King. They are unsheathing their swords. What do you do? Use Igni (fire) or use  Quen (forcefield)? (igni / quen)")
		var iOrQ string
		fmt.Scanln(&iOrQ)

		if iOrQ == "igni" {
			fmt.Println("You cast Igni.")

			if health > 50 {
				fmt.Println("The guards start to burn. You stab each and every one of them... killing them. You then rush into the house to find Ciri. The cage she was locked in is broken, you notice only one pair of footprints which are the size of Ciri's.")
				time.Sleep(2 * time.Second)
				searchCiri()
			} else if health < 50 {
				fmt.Println("The guards start to burn. You try to stab them but one of the guards get up and plunge his steel sword into your chest.")
				os.Exit(0)
			}
		} else if iOrQ == "quen" {
			fmt.Println("You cast Quen.")
			fmt.Println("All the guards try to hit you all at once but your force shield deflects it. You slice them one by one with ease... killing them. You then rush into the house to find Ciri. The cage she was locked in is broken, you notice only on pair of footprints which are the size of Ciri's.")
			time.Sleep(2 * time.Second)
			searchCiri()
		}
	} else if aCOrT == "contract" {
		fmt.Println("You: How about a contract? Any monsters you need gone?")
		time.Sleep(2 * time.Second)
		fmt.Println("King: There are packs of wolves attacking my soldiers relentlessly. Get rid of them and I will give you Ciri.")
		time.Sleep(2 * time.Second)
		fmt.Println("You: Fine.")
		time.Sleep(4 * time.Second)
		fmt.Println("You go outside into the forest side. You use your witcher senses to find the wolves.")
		time.Sleep(2 * time.Second)
		fmt.Println("You find the wolves. What do you do? Use Igni (fire) or wait for their attack? (igni / wait)")
		var iOrW string
		fmt.Scanln(&iOrW)

		if iOrW == "igni" {
			fmt.Println("You burn the wolves and with some swift slices, all the wolves die.")
			time.Sleep(2 * time.Second)
			fmt.Println("When you get back to the castle, the King leads you to where Ciri was. But Ciri isn't there anymore. The cage she was locked in is broken, you notice only one pair of footprints which are the size of Ciri's.")
			time.Sleep(2 * time.Second)
			untraceable()
		} else if iOrW == "wait" {
			fmt.Println("A white wolf comes out from the pack of wolfs, not attacking you. It then walks in the opposite direction as if wanting you to follow.")
			time.Sleep(2 * time.Second)
			fmt.Println("You follow the pack of wolves for a while and then you see a shadowy figure.")
			time.Sleep(2 * time.Second)
			fmt.Println("The wolves stop as if signalling that you have reached.")
			time.Sleep(2 * time.Second)
			fmt.Println("You hear the figure cry out.")
			time.Sleep(2 * time.Second)
			fmt.Println("Figure: Geralt? Geralt!")
			time.Sleep(2 * time.Second)
			ciri()
		}

	} else if aCOrT == "threaten" {
		fmt.Println("You: Fine King. You don't want to talk?")
		time.Sleep(2 * time.Second)
		fmt.Println("You then grab the top of his shirt, lifting him up and hanging him over the well.")
		time.Sleep(2 * time.Second)
		fmt.Println("King: What do you think you are doing Witcher!")
		time.Sleep(2 * time.Second)
		fmt.Println("You: Making you talk.")
		time.Sleep(2 * time.Second)
		fmt.Println("You: Tell me where Ciri is or I drop you down the well. A height well enough to kill you.")
		time.Sleep(2 * time.Second)
		fmt.Println("One of the guards rushes at you but you hear his footsteps and elbow him in the nose.")
		time.Sleep(2 * time.Second)
		fmt.Println("You: You have 10 seconds to start talking before I drop you.")
		time.Sleep(2 * time.Second)
		fmt.Println("King: You wouldn't dare!")
		time.Sleep(2 * time.Second)
		fmt.Println("You: Fine fine. Now you have 5 seconds.")
		time.Sleep(2 * time.Second)
		fmt.Println("King: Bullshit!")
		time.Sleep(2 * time.Second)
		fmt.Println("You: 3 seconds.")
		time.Sleep(2 * time.Second)
		fmt.Println("King: No! You wouldn't!")
		time.Sleep(2 * time.Second)
		fmt.Println("You: 2... 1...")
		time.Sleep(1 * time.Second)
		fmt.Println("King: Fine, I'll tell you, I'll tell you! Just don't kill me please.")
		time.Sleep(2 * time.Second)
		fmt.Println("The King tells you that Ciri is locked up inside the house. She is healthy and fine. Just alone.")
		time.Sleep(2 * time.Second)
		fmt.Println("You rush into the house to find Ciri. The cage she was locked in is broken, you notice only one pair of footprints which are the size of Ciri's.")
		time.Sleep(2 * time.Second)
		searchCiri()
	}
}

func untraceable() {
	fmt.Println("Using your Witcher senses and Ciri's footprints, you try to find Ciri.")
	time.Sleep(2 * time.Second)
	fmt.Println("You see the footprints leave the castle into the forest. You proceed to follow the trail.")
	time.Sleep(2 * time.Second)
	fmt.Println("As you enter the forest, you end up at the same pack of wolves you just killed.")
	time.Sleep(2 * time.Second)
	fmt.Println("The trail has gone cold and you are unable to find Ciri.")
}

func searchCiri() {
	fmt.Println("Using your Witcher senses and Ciri's footprints, you try to find Ciri.")
	time.Sleep(2 * time.Second)
	fmt.Println("You see the footprints leave the castle into the forest. You proceed to follow the trail.")
	time.Sleep(2 * time.Second)
	fmt.Println("As you enter the forest, you hear a pack of wolves.")
	time.Sleep(2 * time.Second)
	fmt.Println("You perpare to kill the wolves by unsheathing your sword but a white wolf comes out from the pack and looks at you. It then walks in the opposite direction as if signalling for you to follow.")
	time.Sleep(2 * time.Second)
	fmt.Println("The wolves stop and you hear someone cry out.")
	time.Sleep(2 * time.Second)
	fmt.Println("Figure: Geralt? Geralt!")
	time.Sleep(2 * time.Second)
	ciri()
}

func ciri() {
	fmt.Println("You: Who are you?")
	time.Sleep(2 * time.Second)
	fmt.Println("Figure: It's me! Ciri!")
	time.Sleep(2 * time.Second)
	fmt.Println("You run over to Ciri.")
	time.Sleep(2 * time.Second)
	fmt.Print("Both of you hug. Lasting around a minute.")
	time.Sleep(2 * time.Second)
	fmt.Println("Then Ciri asks...")
	time.Sleep(2 * time.Second)
	fmt.Println("Ciri: Where's Yennefer?")
	time.Sleep(2 * time.Second)
	fmt.Println("The end.")
}