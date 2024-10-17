package pirateborg

import (
	"fmt"

	"github.com/genrpg/utils"
)

// 1-20 sailor, all else has one entry
var backgrounds = map[string][3]string{
	// Workshop this more, first entry is money, second is items, third is a part of the background, ors are tricky
	"sailor":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a reason to go to sea"},
	"actor":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a play or book", ""},
	"apothecary":         {fmt.Sprintf("%ds", 2*utils.D(6)*10), "healing kit", ""},
	"artist":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "art supplies", ""},
	"assassin":           {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a dagger", ""}, // weapon
	"bandit":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a club", ""},   // weapon
	"barkeep":            {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a flask of fine rum", ""},
	"blacksmith":         {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a set of files and tools", ""},
	"bosun":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "the respect of some crew"},
	"business man":       {fmt.Sprintf("%ds", 3*utils.D(6)*10), "", "a profitable past"},
	"cabin boy":          {fmt.Sprintf("%ds", 2*utils.D(4)*10), "a small toy or trinket", ""},
	"captain":            {fmt.Sprintf("%ds", 2*utils.D(8)*10), "a spyglass", ""},
	"former captive":     {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "your freedom"},
	"castaway":           {fmt.Sprintf("%ds", 2*utils.D(4)*10), "a keepsake from the island", ""},
	"chef":               {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a book of recipes", ""},
	"cook":               {fmt.Sprintf("%ds", 2*utils.D(6)*10), "some fine cooking spices", ""},
	"craftsman":          {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a set of tools", ""},
	"criminal":           {fmt.Sprintf("%ds", utils.D(10)*10), "lockpick or a crowbar", ""},
	"cultist":            {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a book of scripture", ""},
	"deserter":           {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a small bounty on your head"},
	"doctor":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a medical kit", ""},
	"explorer":           {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a map and compass", ""},
	"farmer":             {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a reason to leave your farm"},
	"first mate":         {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a loyal friend"},
	"gambler":            {fmt.Sprintf("%ds", utils.D(12)*10), "", "a sizable debt"},
	"grave robber":       {fmt.Sprintf("%ds", 2*utils.D(6)*10), "something from a corpse" + I("(roll on \"Loot the Body\" table)"), ""}, // special
	"guard":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a keen eye for mischief"},
	"gunner":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "skill at shooting"}, // special
	"harlot":             {fmt.Sprintf("%ds", 2*utils.D(8)*10), "", "an admirer and syphilis"},
	"heretic":            {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a blasphemous disposition"},
	"homemaker":          {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a loved one back home"},
	"hunter":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a hunting knife", ""}, // weapon
	"former servant":     {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "someone looking for you"},
	"innkeeper":          {fmt.Sprintf("%ds", 3*utils.D(4)*10), "", "a small inn somewhere"},
	"landowner":          {fmt.Sprintf("%ds", 3*utils.D(6)*10), "", "property somewhere"},
	"loner":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a dislike of others"},
	"cartographer":       {fmt.Sprintf("%ds", 2*utils.D(6)*10), "map making tools", ""},
	"marine":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a military cutlass", ""}, // weapon
	"medic":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a medical kit", ""},
	"merchant":           {fmt.Sprintf("%ds", 2*utils.D(8)*10), "", "good negotiating skills"}, // special
	"military":           {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "combat training"},         // special
	"missionary":         {fmt.Sprintf("%ds", 2*utils.D(4)*10), "a holy symbol and scripture", ""},
	"monk":               {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a holy symbol and scripture", ""},
	"musician":           {fmt.Sprintf("%ds", 2*utils.D(6)*10), "[roll on instrument table, pg. 27]"}, // special
	"naval deserter":     {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a price on your head"},
	"navigator":          {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a compass", ""},
	"noble":              {fmt.Sprintf("%ds", 3*utils.D(6)*10), "", "good manners"},
	"nobody":             {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a desire for a purpose"},
	"nurse":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a medical kit", ""},
	"officer":            {fmt.Sprintf("%ds", 2*utils.D(8)*10), "a cutlass", "military training"},
	"orphan":             {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a longing for new parental figures"},
	"performer":          {fmt.Sprintf("%ds", 2*utils.D(6)*10), "[roll on instrument table, pg. 27]"}, // special
	"philosopher":        {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a book", "a puzzling disposition"},
	"pilot":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a compass", ""},
	"pirate":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "an empty bottle of rum", ""},
	"politician":         {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "skill at deception"},
	"priest":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a holy symbol and scripture", ""},
	"privateer":          {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a letter of marque", ""},
	"quartermaster":      {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "the respect of a ship's crew"},
	"refugee":            {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "strong survival instincts"}, // special
	"religious follower": {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a strong sense of faith"},
	"revolutionary":      {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "bold plans and a few allies"},
	"rumored sorcerer":   {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "dark and powerful knowledge"},
	"runaway":            {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a desire for a new home"},
	"sail maker":         {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "skill with sails"},
	"scholar":            {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a few books", ""},
	"scoundrel":          {fmt.Sprintf("%ds", 2*utils.D(8)*10), "", "a few enemies along the way"},
	"scout":              {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a map", "keen senses"},
	"shaman":             {fmt.Sprintf("%ds", 2*utils.D(4)*10), "herbs", "mystic knowledge"},
	"shipwright":         {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "skill with woodworking"},
	"smuggler":           {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a bounty on your head"},
	"soldier":            {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a tactical mind"},
	"spy":                {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a dark cloak", ""},
	"student":            {fmt.Sprintf("%ds", 2*utils.D(4)*10), "", "a strong will to learn"},
	"surgeon":            {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a medical kit", ""},
	"thief":              {fmt.Sprintf("%ds", utils.D(12)*10), "lockpicks", ""},
	"vagabond":           {fmt.Sprintf("%ds", 2*utils.D(4)*10), "a leather backpack", ""},
	"victim":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "", "a haunted past"},
	"warrior":            {fmt.Sprintf("%ds", 2*utils.D(4)*10), "a cultural weapon", ""}, // weapon
	"whaler":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "a harpoon", ""},
	"writer":             {fmt.Sprintf("%ds", 2*utils.D(6)*10), "ink, quills, and parchment", ""},
}

var loot = map[string]string{
	"00":    "Roll again twice.",
	"01-20": "The result on the Ancient Relics table (pg. 62).",
	"21":    B("Skull with glowing green sockets."),
	"22":    fmt.Sprintf("%s (d6).", B("A nasty-looking knife")),
	"23":    B("Wanted poster with a picture of one of the PCs on it."),
	"24":    fmt.Sprintf("%s %s", B("Black candle."), I("When lit, its purple flame forms a skull.")),
	"25":    B("Single golden manacle with 5 links."),
	"26":    fmt.Sprintf("%s %s", B("A leather journal"), I("containing 1 sea shanty (pg. 68).")),
	"27":    fmt.Sprintf("%s %s", B("Oil lantern filled with dark green liquid."), I("It burns a pale green light for 5 feet but never runs out.")),
	"28":    B("Fine metal flask."),
	"29":    B("Bag of white powder."),
	"30":    fmt.Sprintf("%s %s", B("Small box with d12 black pearls"), I("worth a fortune.")),
	"31":    B("Sea shell lined with mother of pearl."),
	"32":    fmt.Sprintf("%s %s", B("Pipe carved out of whale bone."), I("A mysterious map is carved in the bowl.")),
	"33":    B("Dead rat."),
	"34":    fmt.Sprintf("%s %s", B("Jewel-encrusted egg."), I("Agility DR18 to open, or it breaks. Clockwork inside is worth 500s.")),
	"35":    fmt.Sprintf("%s %s", B("d8 crab claws."), I("Throwing them before casting a ritual lowers the DR by 1.")),
	"36":    B("Stone ring with an engraved rune."),
	"37":    fmt.Sprintf("%s %s", B("Deep blue gemstone."), I("It sparkles in the moonlight.")),
	"38":    fmt.Sprintf("%s %s", B("Obsidian figurine of a Kraken."), I("The Kraken won’t attack you.")),
	"39":    fmt.Sprintf("%s %s", B("Parrot feather."), I("+1 Devil's Luck each dawn.")),
	"40":    B("Some rotten dried fruit."),
	"41":    fmt.Sprintf("%s %s", B("Recipe for turtle stew."), I("If made, everyone who smells it passes out. At sea, the crew awakens with their ship drifting near a mysterious island.")),
	"42":    B("Small book on tying sailor knots."),
	"43":    fmt.Sprintf("%s %s", B("Deck of playing cards."), I("The queens are mermaids.")),
	"44":    B("Random bomb (pg. 53)."),
	"45":    fmt.Sprintf("%s %s", B("Jar of d10 eyeballs preserved in white rum."), I("They still see.")),
	"46":    B("Set of lock picks."),
	"47":    B("Paper doll painted with blood."),
	"48":    fmt.Sprintf("%s %s", B("Deep green gemstone"), I("worth 100s.")),
	"49":    B("Small wooden flute."),
	"50":    "50s.",
	"51":    "51s.",
	"52":    "52s.",
	"53":    "53s.",
	"54":    "54s.",
	"55":    "55s.",
	"56":    "56s.",
	"57":    "57s.",
	"58":    "58s.",
	"59":    "59s.",
	"60":    fmt.Sprintf("%s %s", B("Leather eye patch."), I("That eye can see in the dark.")),
	"61":    B("Flintlock pistol with tally marks."),
	"62":    B("d4 gold teeth."),
	"63":    B("Letter from a bonnie lass."),
	"64":    fmt.Sprintf("%s %s", B("A ship's schedule"), I("with details of a treasure ship.")),
	"65":    B("Jar of black sand."),
	"66":    fmt.Sprintf("%s %s", B("Book of dark rituals."), I("Test Spirit DR14 to learn one random ritual (pg. 64), or else permanently lose 1 Spirit.")),
	"67":    B("Handwritten collection of ghost stories."),
	"68":    fmt.Sprintf("%s %s", B("Book:"), B(I("A Guide to Sailor Tattoos."))),
	"69":    B("A live rat."),
	"70":    B("d4 dead fish."),
	"71":    fmt.Sprintf("%s %s", B("Vial of blowfish poison."), I("Test Toughness DR12 or take d10 damage.")),
	"72":    B("Treasure map (pg. 119)."),
	"73":    fmt.Sprintf("%s %s", B("A glass eye"), I("that always looks West.")),
	"74":    B("Broken cutlass hilt (d4)."),
	"75":    fmt.Sprintf("%s (3d4), %s", B("Glass dagger"), I("breaks after 1 use.")),
	"76":    fmt.Sprintf("%s %s", B("Book: The Secret Art of Fencing in the Age of Gunpowder."), I("Test Presence DR12 to permanently gain +1 Strength. It only works once.")),
	"77":    B("Bottle filled with iridescent dust."),
	"78":    fmt.Sprintf("%s %s", B("Lint."), I("Flammable.")),
	"79":    fmt.Sprintf("%s %s", B("Golden idol."), I("Priceless, belongs in a museum.")),
	"80":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 80s")),
	"81":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 81s")),
	"82":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 82s")),
	"83":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 83s")),
	"84":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 84s")),
	"85":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 85s")),
	"86":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 86s")),
	"87":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 87s")),
	"88":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 88s")),
	"89":    fmt.Sprintf("%s %s.", B("d8 gold doubloons"), I("worth 89s")),
	"90+":   B(" Random Thing of Importance (pg. 60)."),
}

var thingsOfImportance = []string{"animal pelt", "oyster pearl", "silver locket", "conch shell", "pipe carved from wood", "pipe carved from bone", "small jade figurine", "ancient gold coin", "ruined piece of a treasure map", "map of an unknown place", "diary written by an ancestor", "silver ring", "ivory chess piece", "sea creature carved from obsidian", "spherical prism", "jar containing a severed hand", "necklace of bones & feathers", "book of scripture", "novel you loved as a child", "bizarre silk handkerchief", "pouch containing animal teeth", "old fillet knife", "fossil of an extinct fish", "piece of colorful coral", "small ship in a bottle", "letter from a loved one", "the journal of a dead explorer", "stone embossed with a mermaid", "vial of holy water from clergy in your hometown", "the remains of a small squid in a jar", "precious cooking salts in a tiny chest", "tankard made from a horn", "jar of the finest tobacco", "golden letter opener", "small, cast bronze owl figurine", "collection of sea shells and rocks", "necklace carved from jade", "a recently deceased relative's will naming you as the sole heir", "drawing of a loved one", "bag of “magical” white powder", "old rusted key with a blue gem that glows in the moonlight", "compass that doesn't point north", "clay jar you are using as an urn", "definitive proof of an enemy's (or loved one's) crime", "small golden bell", "old bottle of red wine (Bordeaux, incredible vintage)", "jar of dried jellyfish dust", "multi-colored feather", "necklace from a loved one", "ring that doesn't fit on your fingers", "single diamond earring", "finely made leather eye patch", "set of gardening tools", "dried flower", "animal skull", "human skull", "gem that glows in seawater", "dinosaur or monster bone or claw", "jar of fireflies", "leather-bound tome in a language you don't recognize", "blueprints to a new type of ship", "carved arrowhead", "stone tablet inscribed with ancient pictographs or hieroglyphs", "perfect cube made of crystal", "tattoo, d4: 1 love, 2 revenge, 3 ancestors, 4 unknown origin", "bottle of perfumed oil", "broken set of manacles", "broken compass", "pistol with one shot meant for someone special", "flag of personal significance", "broken spyglass with a scroll or  map hidden inside", "length of rope you made", "carved gaming pieces", "set of rune stones", "twig from a very old tree", "noose taken from a corpse", "6' length of chain", "4d10 scars from lashes on your back", "long scar on your face", "two coconut shells", "dark robe, cape, or cloak", "cask of strong sassafras beer", "set of keys on a large key ring", "small keg of something valuable (rum, powder, ashes, ASH)", "magnifying lens (glass only)", "cork from a bottle, from a special occasion", "cannonball", "deck of cards with 1d4 cards missing and 1d6 'extra' cards", "garment from someone special", "wanted poster, d4: 1 legend, 2 enemy, 3 loved one, 4 stranger", "fancy wig", "letter of political importance", "tanned whale skin or jar of blubber", "petrified egg", "monkey paw extending 1 finger", "memorized poem that sounds  like a map", "medallion that might be the top of a staff", "talisman shaped like a snake", "glass vial of dark blood", "shard of crystal"}

// Endpages, realistically, a GM screen

var meleeWeapons = map[string]any{
	"anchor":            []string{"d10", "60s", "2 handed"},
	"bayonet":           []string{"d4", "15s"},
	"belaying pin":      []string{"d4", "10s"},
	"boarding axe":      []string{"d6", "20s"},
	"boarding pike":     []string{"d10", "60s", "2 handed"},
	"broadsword":        []string{"d8", "35s"},
	"broken bottle":     []string{"d2", "-"},
	"cat o' nine tails": []string{"d4", "15s"},
	"chain":             []string{"d6", "25s"},
	"cudgel":            []string{"d4", "10s"},
	"cutlass":           []string{"d6", "25s"},
	"fine rapier":       []string{"d8", "50s"},
	"grappling hook":    []string{"d6", "35s"},
	"harpoon":           []string{"d8", "35s"},
	"hatchet":           []string{"d6", "20s"},
	"heavy club":        []string{"d6", "20s"},
	"hook":              []string{"d4", "8s"},
	"knife/dagger":      []string{"d4", "10s"},
	"machete":           []string{"d6", "25s"},
	"marlinspike":       []string{"d4", "10s"},
	"officer’s cutlass": []string{"d8", "50s"},
	"rapier":            []string{"d6", "30s"},
	"scimitar":          []string{"d6", "25s"},
	"smallsword":        []string{"d4", "20s"},
	"tomahawk":          []string{"d6", "20s"},
	"unarmed/insults":   []string{"d2", "-"},
	"whale bone":        []string{"d4", "-"},
	"wood plank":        []string{"d4", "-"},
}

var rangedWeapons = map[string]any{
	"blowpipe":         []string{"–", "30s", I("See darts")},
	"blunderbuss":      []string{"d4(d10)", "65s", "Reload 2, d10 within 10'"},
	"buccaneer musket": []string{"2d8", "100s", "Reload 2"},
	"flintlock pistol": []string{"2d4", "50s", "Reload 2"},
	"harpoon gun":      []string{"d8", "60s", "Strength DR12 or pulled"},
	"musket":           []string{"2d6", "80s", "Reload 2"},
	"throwing axes":    []string{"d6", "20s "},
	"throwing knives":  []string{"d4", "8s "},
}

var ammo = map[string]any{
	"20 rounds of shot": []string{"-", "10s"},
	"10 berserk darts":  []string{"d4+", "20s", "Toughness DR12 or attack closest creature for d4 rounds"},
	"10 poison darts":   []string{"d4+", "20s", "Toughness DR12 or d6 damage"},
	"10 sleep darts":    []string{"d4+", "20s", "Toughness DR12 or sleep d6 rounds"},
}

var bombs = map[string]any{
	"smoke bomb":         []string{"-", "10s", "Blind for d4 rounds"},
	"improvised grenade": []string{"d10", "20s"},
	"clay grenade":       []string{"2d8", "30s"},
	"iron grenade":       []string{"3d6", "40s"},
	"fire pot":           []string{"d6", "15s", "d6/turn: 1-2 spreads, 6 fire goes out."},
	"stink ball":         []string{"2d4", "20s", "Toughness DR12 or poisoned"},
}

var bombsText = "Test Agility DR12 to hit an area. Hit: Creatures within 5' of the area test Agility DR12 or take damage. Fumble: hit self and/or d4 allies instead. Crit: x2 dmg."

var ancientRelics = map[int]Item{
	1:  {Name: "Cross of the Paragon", Extra: "One ally gets " + B("+1") + " to attack and " + B("+1") + " to damage for " + B("d6") + " turns."},
	2:  {Name: "Conch Shell from the Abyss", Extra: "Ask a nearby corpse (or any creature that died at sea within 100 miles) one question"},
	3:  {Name: "Map Inked in Ectoplasm", Extra: "Learn the location of all traps and secret doors within 30' for " + B("d4+SPIRIT") + " rounds"},
	4:  {Name: "Will-o'-the-Wisp Lantern", Extra: "Emit 15' of light or darkness for " + B("d6+SPIRIT") + " rounds."},
	5:  {Name: "Pages from the Necronomicon", Extra: "All creatures that can hear your voice test DR14 or lose " + B("d4+SPIRIT") + " HP (ignore armor)."},
	6:  {Name: "Rune Encrusted Flintlock Pistol", Extra: "One creature you see loses " + B("d6+SPIRIT") + " HP (ignore armor). Takes 1 action to reload."},
	7:  {Name: "Jade Die", Extra: "Roll a die. Odd: you gain " + B("d8") + " temporary HP. Even: Choose a creature. It gets " + B("+d8") + " on its next damage roll."},
	8:  {Name: "Undead Bird", Extra: "It can speak with animals (dead or alive) for " + B("d6+SPIRIT") + " rounds."},
	9:  {Name: "Mermaid Scales", Extra: "Eat a scale: breathe underwater for " + B("d4") + " hours."},
	10: {Name: "Charon's Obol", Extra: "If you are killed, return to life the next round with 1 HP. Disappears after one use."},
	11: {Name: "Cup of the Carpenter", Extra: "Choose a creature to regain " + B("d6+SPIRIT") + " HP."},
	12: {Name: "Heart of the Sea", Extra: "Create or destroy 15 gallons water or 30 square feet of fog."},
	13: {Name: "Necklace of Eyeballs", Extra: "Become invisible for " + B("d6+SPIRIT") + " rounds or until you attack or take damage. Attack and defend with " + B("DR6") + "."},
	14: {Name: "Crown of the Sunken Lord", Extra: "A water shield surrounds you. " + B("-d2") + " protection for " + B("d2+SPIRIT") + " rounds (in addition to armor)."},
	15: {Name: "Crystalline Skull", Extra: "The skull can hear & repeat the thoughts of a nearby creature for " + B("d6+SPIRIT") + " minutes."},
	16: {Name: "Codex Tablet", Extra: "Read and understand any language, glyphs, or runes for " + B("1+SPIRIT") + " rounds."},
	17: {Name: "Skeleton Key", Extra: "Open any door or lock. Crumbles after use."},
	18: {Name: "Mummified Monkey Head", Extra: "The head speaks: 1 creature tests SPIRIT DR12 or must obey a 1 word command."},
	19: {Name: "Great Old One Figurine", Extra: "One human is terrorized for " + B("d4") + " rounds unless they succeed a PRESENCE DR14 test. They can test each round."},
	20: {Name: "Broken Compass", Extra: "The compass points in the direction of an object you know of for " + B("1+SPIRIT") + " rounds."},
}

var distinctiveFlaws = []string{"Drunken lush", "Stubborn", "Mocking sardonic cheer", "Way too loud", "Stupid", "Coward", "Cocky", "Slightly deranged", "Aggressive", "Anxious", "Cheater", "Selfish", "Lazy", "Hedonistic", "Impulsive", "Ostentatious", "Paranoid", "Pretentious", "Sadistic", "Disloyal"}

var physicalTrademarks = []string{"Cursed: visibly part skeleton/ghost/water/flames/coral", "Missing an eye", "Matted, dreaded hair", "Missing a leg: pegleg or crutch", "Missing a hand: hook or claw instead.", "Missing an ear", "Many, many tattoos", "Never blinks. Ever.", "Rotten or broken teeth", "Twitches constantly, especially trigger finger", "A nigh incurable case of scurvy: permanently bleeding gums", "Infested with bugs", "Gnarly facial scar", "Hideously ugly", "Corpulent", "Increasingly gangrenous", "Putrid, bilge stench", "Contagious", "Gaunt & frail", "So good looking people are " + I("jealous")}

var idiosyncrasies = []string{"You smoke " + I("constantly") + ", and cough even more.", "\"Functioning\" alcoholic. You're probably drunk right now.", "You bet on " + I("everything") + " possible.", "Constantly counting. Teeth, cannon balls... " + I("everything") + ".", "Rats are your favorite meal.", "You know every tall tale ever told. You make sure everyone else knows you know them.", "You are afraid of prime numbers larger than 3. d20 rolls of 5, 7, 11, 13, and 17 fill you with superstitious terror.", "You become a murderous grump when hungry.", "Habitual procrastinator... if you even finish the task.", "You are a voluntary insomniac. Sleep is for the dead.", "You prefer to shoot first and never ask questions.", "Overly, annoyingly religious.", "You collect something, and you often talk to your collection. They are your " + I("friends") + ".", "Always trying to trick your crewmates, just for fun.", "Why pay for anything when you can steal it?", "You talk to yourself when alone, but you often think you are alone when you aren't.", "You secretly enjoy the taste of human flesh.", "You always say you \"know the right way\" but are prone to getting lost.", "You blame everyone but yourself for all of your mistakes.", "Extremely obsessive with tasks and relationships."}

var unfortunateIncidents = []string{"Your loved ones were burned alive. Revenge is imminent.", "You are a " + I("known") + " pirate. You face the gallows if caught.", "You betrayed former crewmates. Now they hunt " + I("you") + ".", "You were marooned on an island for far too long. The voices " + I("must") + " be real.", "You stole a ship. The owner wants your money or your head, but will settle for both.", "You escaped captivity, and will " + I("never") + " go back.", "A close relative has become your greatest enemy.", "The last three ships you crewed all sank.", "Your last crew was killed by undead. They left you alive on purpose.", "Political leaders hold your loved one(s) captive.", "An undead spirit you don't like possesses you regularly.", "You wronged an infamous pirate lord.", "You narrowly escaped a cannibalistic ending, but you didn't escape " + I("that smell") + ".", "You slaughtered them. Like " + I("animals") + ".", "You are the mysterious lone survivor of a treasure expedition gone awry.", "[d2] 1: Failed mutineer. 2: Successful mutineer.", "A silent ghost haunts you. It is always there, but only you can see it.", "You deserted the military, but you're not sure who knows so.", "You have no memory before a few days ago.", "You died once already, but Hell didn't want you."}

// TODOs:
//	Endpages content
//	The rest of it

// d10
var vesselClass = []string{"raft", "longboat", "tartane", "sloop", "brigantine", "fluyt", "frigate", "galleon", "man-of-war", "ship of the line"}

// d6
var armament = []string{"Merchant: No weapons.", "Lightly armed: reduce damage die size by one.", "Normal armament", "Warship: Double broadside"}

// d6
var crewQuantity = []string{"Short-handed: half as many.", "Standard crew.", "Ready for war: twice as many.", "Ready to raid: as many as possible."}

// 2d6
var crewQuality = []string{"Near mutiny and/or untrained.", "Miserable and/or novice.", "Average.", "Fresh from shore leave and/or experienced.", "Prosperous/loyal and/or military training"}

// d66(d36)
var shipName = []string{"Banshee's Wail", "Revenant", "Void Ripper", "Mermaid's Tear", "Carrion Crow", "Executioner's Hand", "Poseidon's Rage", "Adventure's Ghost", "Widow's Revenge", "Blood Moon", "Devil's Scorn", "Gilded Parrot", "Monolith", "Black Tide", "Raven's Wrath", "Coral Corsair", "Hellspire", "Vendetta", "Crimson Tempest", "Royal Tomb", "Guillotine", "Neptune's Maiden", "Cadaver's Call", "Heart of the Sea", "Demonwake", "Bride of the Abyss", "Annihilation", "Golden Glaive", "Necrobile", "Grave Witch", "Loa's Lament", "Hunsi Kanzo", "Dragon from the Deep", "Leviathan's Flood", "Kraken's Maw", "Harlequin's Curse"}

// d12 - roll further
var mundaneCargo = []string{"food or crops , 250s", "spices or oils, 350s", "trade goods, 400s", "livestock, 400s", "sugar, 500s", "rum, 1000s", "munitions, 2000s", "tobacco, 1000s", "wine, 2000s", "antiques, 2000s", "lumber, 1000s", "special cargo"}

// d12 - roll further
var specialCargo = []string{"raw silver ore, 5000s", "golden coins and treasures, 10000s", "religious leader(s)", "important prisoner(s)", "political or military figure(s)", "relics or a rare artifact, 4000s", "sea monster bones, 2500s", "exotic animals, 2000s", "d10 locked chests, 2d8 x 100s each", "d20 crates of ASH , see pg. 25", "imprisoned undead", "a sorcerer with a tome of d4 Arcane Rituals (pg. 64)"}

// d8
var optionalVesselPlotTwist = []string{"Deadly disease on board.", "Crew are impostors.", "Crew is mute.", "The PCs know this crew.", "Everyone on board was thought to be dead.", "Ghost ship.", "They're all zombies.", "Someone on board is related to a PC's backstory."}

// d12
var whereIsIt = []string{"shallow waters (submerged at high tide)", "suspended on shoals (half underwater)", "beach wrecked", "adrift at sea", "anchored off the coast", "suspended on rocks or a coastal cliff", "up a dried-up riverbed", "in the middle of the jungle, forest, or desert", "drifting into port", "orbiting a maelstrom", "floating the waters that lead to the underworld", "in the nightmares of cursed sailors"}

// d12
var typeOfShip = []string{"several ships fastened together (roll d2 more times)", "sloop", "tartane", "giant jury-rigged raft", "brigantine", "frigate", "ancient vessel", "galleon", "fluyt", "man-of-war", "ship of the line", "otherworldly"}

// d8
var whatHappenedHere = []string{"demolished during a storm or hurricane", "abandoned for an unknown reason, mostly intact", "ripped in two by a monster from the deep", "run aground or scrapped some shoals", "raided by blood-thirsty undead", "destroyed in naval combat", "wrecked in foggy conditions", "mutiny fueled blood bath"}

// d8
var inOneOfTheRooms = []string{"filled with large eggs", "bodies of former crew, freshly dead, terror on their faces", "indecipherable glyphs carved into the wood beams", "water damaged books and letters", "hundreds of eyeballs hanging from strings", "a gaping hole in the hull", "rotting food and animal corpses", "glass bottles filled with [d6]: 1 rum 2 potions (pg. 70) 3 fortified wine 4 blood 5 holy water 6 excrement (human?)"}

// d12
var oddFeature = []string{"walls and floor covered in coral and barnacles", "piles of bleached white bones throughout", "hundreds of small crabs nests", "cargo hold is filled with 6\" of blood", "charred wood and fire damage", "faintly glows in the dark", "mysterious slime covers most surfaces", "ornately decorated in gold leaf and velvet", "a thick layer of ash coats everything", "gravity behaves as if underwater", "signs of torture, sacrifice, and blood-letting", "bioluminescent plants bloom from hull at night"}

// d8 // TODO: Derelict Ships
var development = []string{}
