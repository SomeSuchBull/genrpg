package pirateborg

import (
	"fmt"

	"github.com/genrpg/utils"
)

// d100; 1-20 sailor, all else has one entry
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

// d100
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

// d100
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

// d20
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

// d20
var distinctiveFlaws = []string{"Drunken lush", "Stubborn", "Mocking sardonic cheer", "Way too loud", "Stupid", "Coward", "Cocky", "Slightly deranged", "Aggressive", "Anxious", "Cheater", "Selfish", "Lazy", "Hedonistic", "Impulsive", "Ostentatious", "Paranoid", "Pretentious", "Sadistic", "Disloyal"}

// d20
var physicalTrademarks = []string{"Cursed: visibly part skeleton/ghost/water/flames/coral", "Missing an eye", "Matted, dreaded hair", "Missing a leg: pegleg or crutch", "Missing a hand: hook or claw instead.", "Missing an ear", "Many, many tattoos", "Never blinks. Ever.", "Rotten or broken teeth", "Twitches constantly, especially trigger finger", "A nigh incurable case of scurvy: permanently bleeding gums", "Infested with bugs", "Gnarly facial scar", "Hideously ugly", "Corpulent", "Increasingly gangrenous", "Putrid, bilge stench", "Contagious", "Gaunt & frail", "So good looking people are " + I("jealous")}

// d20
var idiosyncrasies = []string{"You smoke " + I("constantly") + ", and cough even more.", "\"Functioning\" alcoholic. You're probably drunk right now.", "You bet on " + I("everything") + " possible.", "Constantly counting. Teeth, cannon balls... " + I("everything") + ".", "Rats are your favorite meal.", "You know every tall tale ever told. You make sure everyone else knows you know them.", "You are afraid of prime numbers larger than 3. d20 rolls of 5, 7, 11, 13, and 17 fill you with superstitious terror.", "You become a murderous grump when hungry.", "Habitual procrastinator... if you even finish the task.", "You are a voluntary insomniac. Sleep is for the dead.", "You prefer to shoot first and never ask questions.", "Overly, annoyingly religious.", "You collect something, and you often talk to your collection. They are your " + I("friends") + ".", "Always trying to trick your crewmates, just for fun.", "Why pay for anything when you can steal it?", "You talk to yourself when alone, but you often think you are alone when you aren't.", "You secretly enjoy the taste of human flesh.", "You always say you \"know the right way\" but are prone to getting lost.", "You blame everyone but yourself for all of your mistakes.", "Extremely obsessive with tasks and relationships."}

// d20
var unfortunateIncidents = []string{"Your loved ones were burned alive. Revenge is imminent.", "You are a " + I("known") + " pirate. You face the gallows if caught.", "You betrayed former crewmates. Now they hunt " + I("you") + ".", "You were marooned on an island for far too long. The voices " + I("must") + " be real.", "You stole a ship. The owner wants your money or your head, but will settle for both.", "You escaped captivity, and will " + I("never") + " go back.", "A close relative has become your greatest enemy.", "The last three ships you crewed all sank.", "Your last crew was killed by undead. They left you alive on purpose.", "Political leaders hold your loved one(s) captive.", "An undead spirit you don't like possesses you regularly.", "You wronged an infamous pirate lord.", "You narrowly escaped a cannibalistic ending, but you didn't escape " + I("that smell") + ".", "You slaughtered them. Like " + I("animals") + ".", "You are the mysterious lone survivor of a treasure expedition gone awry.", "[d2] 1: Failed mutineer. 2: Successful mutineer.", "A silent ghost haunts you. It is always there, but only you can see it.", "You deserted the military, but you're not sure who knows so.", "You have no memory before a few days ago.", "You died once already, but Hell didn't want you."}

// TODOs:
//	Endpages content
//	The rest of it

// VESSEL --

// d10
var vesselClass = []string{"raft", "longboat", "tartane", "sloop", "brigantine", "fluyt", "frigate", "galleon", "man-of-war", "ship of the line"}

// d10
var vesselClassCargo = []int{0, 0, 1, 2, 3, 5, 4, 6, 4, 4}

// d6
var vesselArmament = []string{"Merchant: No weapons.", "Lightly armed: reduce damage die size by one.", "Normal armament", "Normal armament", "Normal armament", "Warship: Double broadside"}

// d6
var vesselCrewQuantity = []string{"Short-handed: half as many.", "Standard crew.", "Standard crew.", "Ready for war: twice as many.", "Ready for war: twice as many.", "Ready to raid: as many as possible."}

// 2d6
var vesselCrewQuality = map[int]string{
	2:  "Near mutiny and/or untrained.",
	3:  "Near mutiny and/or untrained.",
	4:  "Miserable and/or novice.",
	5:  "Miserable and/or novice.",
	6:  "Miserable and/or novice.",
	7:  "Average.",
	8:  "Average.",
	9:  "Fresh from shore leave and/or experienced.",
	10: "Fresh from shore leave and/or experienced.",
	11: "Prosperous/loyal and/or military training",
	12: "Prosperous/loyal and/or military training",
}

// d66(d36)
var vesselShipName = []string{"Banshee's Wail", "Revenant", "Void Ripper", "Mermaid's Tear", "Carrion Crow", "Executioner's Hand", "Poseidon's Rage", "Adventure's Ghost", "Widow's Revenge", "Blood Moon", "Devil's Scorn", "Gilded Parrot", "Monolith", "Black Tide", "Raven's Wrath", "Coral Corsair", "Hellspire", "Vendetta", "Crimson Tempest", "Royal Tomb", "Guillotine", "Neptune's Maiden", "Cadaver's Call", "Heart of the Sea", "Demonwake", "Bride of the Abyss", "Annihilation", "Golden Glaive", "Necrobile", "Grave Witch", "Loa's Lament", "Hunsi Kanzo", "Dragon from the Deep", "Leviathan's Flood", "Kraken's Maw", "Harlequin's Curse"}

// d12 - roll further
var vesselMundaneCargo = []string{"food or crops, 250s", "spices or oils, 350s", "trade goods, 400s", "livestock, 400s", "sugar, 500s", "rum, 1000s", "munitions, 2000s", "tobacco, 1000s", "wine, 2000s", "antiques, 2000s", "lumber, 1000s", "special cargo"}

// d12
var vesselSpecialCargo = []string{"raw silver ore, 5000s", "golden coins and treasures, 10000s", "religious leader(s)", "important prisoner(s)", "political or military figure(s)", "relics or a rare artifact, 4000s", "sea monster bones, 2500s", "exotic animals, 2000s", "d10 locked chests, 2d8 x 100s each", "d20 crates of ASH , see pg. 25", "imprisoned undead", "a sorcerer with a tome of d4 Arcane Rituals (pg. 64)"}

// d8
var vesselOptionalPlotTwist = []string{"Deadly disease on board.", "Crew are impostors.", "Crew is mute.", "The PCs know this crew.", "Everyone on board was thought to be dead.", "Ghost ship.", "They're all zombies.", "Someone on board is related to a PC's backstory."}

// -- VESSEL

// DERELICT SHIP --

// d12
var derelictShipWhereIsIt = []string{"shallow waters (submerged at high tide)", "suspended on shoals (half underwater)", "beach wrecked", "adrift at sea", "anchored off the coast", "suspended on rocks or a coastal cliff", "up a dried-up riverbed", "in the middle of the jungle, forest, or desert", "drifting into port", "orbiting a maelstrom", "floating the waters that lead to the underworld", "in the nightmares of cursed sailors"}

// d12
var derelictShipTypeOfShip = []string{"several ships fastened together (roll d2 more times)", "sloop", "tartane", "giant jury-rigged raft", "brigantine", "frigate", "ancient vessel", "galleon", "fluyt", "man-of-war", "ship of the line", "otherworldly"}

// d8
var derelictShipWhatHappenedHere = []string{"demolished during a storm or hurricane", "abandoned for an unknown reason, mostly intact", "ripped in two by a monster from the deep", "run aground or scrapped some shoals", "raided by blood-thirsty undead", "destroyed in naval combat", "wrecked in foggy conditions", "mutiny fueled blood bath"}

// d8
var derelictShipInOneOfTheRooms = []string{"filled with large eggs", "bodies of former crew, freshly dead, terror on their faces", "indecipherable glyphs carved into the wood beams", "water damaged books and letters", "hundreds of eyeballs hanging from strings", "a gaping hole in the hull", "rotting food and animal corpses", "glass bottles filled with [d6]: 1 rum 2 potions (pg. 70) 3 fortified wine 4 blood 5 holy water 6 excrement (human?)"}

// d12
var derelictShipOddFeature = []string{"walls and floor covered in coral and barnacles", "piles of bleached white bones throughout", "hundreds of small crabs nests", "cargo hold is filled with 6\" of blood", "charred wood and fire damage", "faintly glows in the dark", "mysterious slime covers most surfaces", "ornately decorated in gold leaf and velvet", "a thick layer of ash coats everything", "gravity behaves as if underwater", "signs of torture, sacrifice, and blood-letting", "bioluminescent plants bloom from hull at night"}

// d8
var derelictShipDevelopment = []string{"Sails! Ship spotted.", "Derelict starts to sink or collapse.", "Adversaries arrive: other pirates, monsters, the military.", "Explosion or fire.", "Extreme weather.", "Thing from the deep complicates things.", "Undead floating nearby.", "A stowaway emerges."}

// d12
var derelictShipCurrentOccupant = []string{"Ghosts of dead sailors. Their speech sounds like whale song.", "Poisonous necro-plague rats, their mother the size of a wolf.", "Dormant zombies waiting for fresh brains, their smell detectable for miles.", "A starving crew of pirate-turned-cannibals, some with freshly missing limbs.", "Skeleton marauders celebrating a recent raid. The captain is an asshole.", "Cultists, chanting, summoning Great Old Ones via human sacrifice.", "3 beautiful sirens. They lust for pearls and have a hoard hidden nearby.", "A tribe of intelligent goblin-like monkeys. They love shiny things and rum.", "A necromancer living in exile with a pet undead jaguar or shark.", "Deep Ones, lead by a shaman, scouting for valuables. Her magic requires water.", "Hundreds of carrion gulls, feasting on a corpse hanging from a mast.", "Dozens of blood drained corpses. A weakened vampire hides below deck."}

// d20
var derelictShipOriginalCargo = []string{"tea leaves or spices", "art or antiques", "coffee beans", "sugar or sugar cane", "rum or wine", "weapons and ammo", "tobacco or cocoa", "exotic creatures", "rice or crops", "black powder [d100 barrels]", "contraband", "textiles", "sacrificial victims", "prisoners", "lumber", "livestock", "a large statue", "hoards of treasure", "native medicine", "a cursed relic"}

// d6
var derelictShipCargoCondition = []string{"missing or ruined", "salvageable (x1/2)", "salvageable (x1/2)", "good (x1)", "good (x1)", "rare and valuable (x2)"}

// -- DERELICT SHIP

// d100
var buriedTreasure = []string{"The Black Spot on a folded piece of paper. Whoever unfolds it dies in d8 days (unless they find a way to lift the curse).", "A Spanish gold treasure hoard worth 50,000s. They want it back NOW.", "3 flags: British, Spanish, French. One torn, one blood spattered, one outdated.", "Freshly decapitated head. It looks exactly like one of the PCs. In its mouth is a pearl the size of a plum.", "Emerald the size of a coconut.", "6 stylized animal heads carved from wood, one on top of the next.", "7 rare gold coins worth d100 silver each.", "Each player rolls all of the dice they have with them: the total in silver.", "An enormous dinosaur bone.", "2 black flintlocks, the name \"Blackwood\" engraved on both. DR -1 to hit.", "A bag of marbles.", "Thick leather belt, an iron kraken buckle. Wearer gets +1 strength.", "A pure white flintlock pistol. It crits on 18+ but shatters on a fumble.", "Half of a metal medallion.", "15 dead men on a dead man's chest. The crew learns 1 random Shanty (pg. 68).", "An elaborate brass anchor.", "An ancient circular stone calendar. Its last entry is the day after tomorrow...", "5,000 gold coins (they are painted wood).", "d6 black candles. Burning them makes all other light sources dim by half.", "d20 silver bars (50s each).", "8' marble statue of a sea monster with rubies for eyes.", "Detailed journal with drawings and descriptions of the local flora and fauna.", "Golden crown and scepter.", "a flute carved from a femur.", "d6 bladders from some sea creature. Each contains a potion (pg. 70).", "Jet black sphere.", "Mother of pearl encrusted lobster shell.", "Black cannon in the shape of a dragon. A ship with it always deals +1 damage.", "Book written in an ancient script. It contains the best recipes known to man.", "300 dead monkeys.", "Dead pirate captain. There is a 2,000s reward for his head.", "32 golden teeth and a set of silver pliers.", "Skeleton of a small animal.", "Single golden arrow.", "35' of hempen rope that will (almost) never break.", "Small sandstone pyramid. Smashing it reveals a black diamond worth 2,000s.", "Mesoan coin on a necklace.", "703 pieces of eight in a wooden crate.", "Dark gemstones. Looking into them reveals an underwater location.", "Ornate model frigate inside a bottle.", "Large fang the size of a femur.", "A fine wooden armoire. Inside are d12 outfits fit for royalty.", "Iron scarab. On sand it explodes for d6 damage, underwater it swims to treasure.", "Large jar filled with 2d20 dried seahorses. Eating one recovers 1 HP. Don't eat too many...", "Pair of finely crafted boots. +1 agility.", "Bloody leather bag with 3 silver rings, 7 stone rings, and 9 iron rings.", "Polished obsidian disc. If you look at the sun through it you immediately gain experience, then it shatters.", "4' x 8' stone slab covered in glyphs. It depicts a woman entering a portal.", "Carved ivory sailing ship.", "Sarcophagus (5,000s). Inside is a scepter that turns human ashes into silver dust.", "Dark green cloak. +1 presence.", "d20 bottles of the finest rum ever made.", "Barrel-sized egg. The inside stirs...", "Taxidermy dog. In its belly is a severed hand holding a set of keys.", "Severed finger. If rested on a map it will point to a lost city of gold.", "d4 barrels filled with blood.", "Smoking pipe shaped like a galleon.", "Hollow cannonball. Inside: a map.", "Clay dagger with glyphs carved in the blade. Read them: learn one random ritual (pg. 64) then it crumbles.", "Chest with 4 different colored vials.", "A figurehead of a gorgeous woman. If attached to a ship, it comes alive and screams in terror.", "Clay jar filled with 13 snakes.", "A skeleton key, pg. 63.", "Locked iron box. Inside is a rotten head that can speak. It knows the way to Hell.", "Set of 6 jade goblets.", "Everything within 66' dies immediately. (Optional: all PCs gain the Haunted Soul: Skeleton class until the treasure is reburied.)", "Silver bracelet of a snake eating its tail.", "Handheld crystal mirror.", "3d20 large pearls.", "Collection of exotic bird feathers (300s).", "Shriveled heart. It starts beating if submerged in sea water.", "Wood coffin. Inside is (d4): 1 a skeleton 2 a mummy 3 a vampire 4 1,000s.", "Map of the Dark Caribbean. A trail of blood appears along your path.", "Jaguar skull made of diamond.", "Key made of water. Glows if submerged.", "Jade gecko wrapped around a coconut sized granite globe.", "Vial of water. Drinking it secretly gives d100 HP.", "Satchel of hauntingly poetic love letters.", "Chess set carved from onyx and jade.", "Wood tube with d6 scrolls, each with a different Ritual (pg. 64). They always succeed but burn up once used.", "Stone tankard, covered with Mesoan glyphs. It never runs out of ale.", "Golden pegleg.", "Gold ship's bell. Ringing it changes the winds direction.", "Plain gold ring with no signs of wear (in fire it projects a clue, pg. 121).", "A chest inside a chest. Roll again.", "Necklace of bones, each with a different engraved rune. +1 toughness.", "Gray, pointy hat. Wearer gets +1 spirit.", "Trident etched with antediluvian runes.", "Bronze skull with d2 gems for eyes.", "Bullwhip. Hits on dr8, d6 dmg, 10' range.", "Amulet of the Sun. Elaborately carved, Ancient, and extremely valuable.", "A large crate. d100 skulls are inside.", "Volcanic glass sculpture of an island. Holding it to the sun reveals a small ruby inside. It's a map.", "Spyglass. It seems broken, but actually peers 24 hours into the past.", "Book of 7 maps in 7 different languages. Each is of a different small island.", "3 turtles shells of different colors. Each is filled with gems worth d100 silver.", "Go. A bag of black stones, a bag of white stones, and a bamboo game board.", "Oil lamp. (The genie inside will grant one wish, but in the worst way possible).", "A wooden box with 24 stacks of 20 gold coins (500s each) and 6 loaded flintlocks.", "1 million pieces of eight. Good luck."}

// TREASURE MAP --

// d4
var treasureMapTypeOfMap = []string{"Simple. A geographic area, X marks the spot.", "Point Crawl. The map shows several icons (pg. 120)icons (pg. 120), but they must be traveled to sequentially to find the next one.", "Scavenger Hunt. The map has only one point of interest or one clue, but finding it reveals more maps or clues.", "Arcane. The map only reveals the next step when the previous one is completed."}

// d12
var treasureMapMaterial = []string{"Parchment paper.", "Parchment paper.", "Parchment paper.", "Parchment paper.", "In an old book.", "Only in someone's mind or memories.", "Burnt into leather.", "Whittled into a bone or skull.", "Minted on a unique coin.", "Embossed into a sea turtle shell.", "Scratched into driftwood.", "Etched into a glass bottle."}

// d8
var treasureMapGeography = []string{"Island", "Coast", "Peninsula", "Cove", "Archipelago", "River", "Lake", "Inland"}

// d6
var treasureMapStarts = []string{"at a totally unknown location.", "nearby.", "within a mile.", "a day's walk away.", "a day's sail away.", "near a well known place."}

// d10
var treasureMapInvisibleInkRevealedBy = []string{"with heat or fire.", "under star or moon light.", "during an eclipse.", "when underwater.", "if soaked in alcohol.", "with magic or a ritual.", "when you speak the password.", "when near a key location or object.", "during the solstice or equinox.", "at a certain time of day."}

// d8
var treasureMapComplication = []string{"Another crew is there when the PCs arrive.", "Monsters guard it.", "It has already been found or is missing.", "It's cursed.", "There are hundreds of corpses.", "The PCs become incredibly lost on the way.", "The whole thing is a setup.", "There is so much treasure everyone becomes murderously paranoid."}

// d20
var treasureMapIcons = []string{"rope bridge", "waterfall", "volcano caldera", "skull rock", "wicked tree", "jungle", "beach caves", "whale graveyard", "white sand beach", "lagoon", "ancient ruins", "burnt down cabin", "swamp", "shipwreck", "cliff side", "large hill", "quicksand", "mountain's peak", "highest point", "lowest point"}

// d10
var treasureMapClueContainer = []string{"note in a scroll case", "small stone, instructions carved underneath", "coconut shell", "message in a bottle", "metal plaque", "letter in a small box", "etched cannon ball", "burnt or carved into wood", "dried blood on a skull", "A potion. The drinker \"knows\" the next clue, but toughness dr12 or becomes poisoned."}

// d20
var treasureMapMarkers = []string{"large boulder", "dead man's bones", "makeshift grave", "old anchor", "ruined rowboat", "lone palm tree", "ship's wheel", "painted rock", "cannonball", "pile of skulls", "volcanic glass", "rusty cutlass", "broken oar", "broken barrel", "ruined cannon", "tallest tree", "strange sea shell", "rotting crate", "stone statue", "tombstone"}

// d12
var treasureMapRiddles = []string{"Arrive at the [icon] but you won't see the [marker] in the light, travel [2d6] paces to the [direction] and wait until day turns to night.\n" + I("The marker only appears in the dark."), "Set your course for the [icon] and ready a shanty, for singing near the [marker] will turn \"none\" into \"plenty\".\n" + I("The treasure or clue appears once music is played."), "Find the [marker] by the [icon], ye soothsayers and witches, for only magic or prayers can lead to these riches.\n" + I("Using a Relic, Ritual, Spell, or Prayer magically reveals the clue or treasure."), "Look not high for salvation, but low for an X, crossed trees you will find, but beware of the hex.\n" + I("Two fallen trees cover the clue or treasure; it's booby trapped."), "In shallow waters near [icon] the [marker] rests under waves, it points in the direction of nearby caves.\n" + I("The marker has a carved or painted arrow that points to a nearby cave with the clue/treasure."), "Head [d8 x 10] paces [direction] of the [icon] then dig a hole, under the [marker] lies the key to your goal.\n" + I("Dig for the clue/treasure."), "Make for the [icon] before the sun hits the sea, the secret lies between the [marker] and a tree.\n" + I("Can't be found during the night."), "Only dead men know this hidden locale, for of all who've seen the [marker] by the [icon] none have lived to tell.\n" + I("Tragedy, a trap, or both haunt this location."), "Near [icon] at midnight the dead rise from their graves, but the one near the [marker] is a crooked-nosed knave.\n" + I("Harmless ghosts rise here at midnight, but one near the marker knows the location of the clue/treasure. He will only reveal it if his body, hanging from a nearby noose, is laid to rest."), "Travel [direction] for 2d2 miles, find the [icon] then rest, only candlelight will reveal the [marker] over the chest.\n" + I("The marker only appears in candle light. The clue/treasure is under it."), "From the [icon] at low tide you will be able to see, the [marker] just [direction] of a large fruit tree.\n" + I("The clue/treasure is buried in the shadow of the tree."), "[Direction] of the [icon] you will find a sight rarely seen, find the [marker] that's hidden under a blanket of green.\n" + I("The marker is hidden under foliage, moss, or grass growing in an odd place.")}

// -- TREASURE MAP

// ISLAND --

// d6
var islandSize = []string{"Tiny", "Small", "Medium", "Large", "Large", "Huge"}

// 2d6
var islandTerrain = map[int]string{2: "An atoll", 3: "It's two islands", 4: "Gentle hills", 5: "Flat with a mountainous ridge", 6: "Mostly flat", 7: "Mostly flat", 8: "Mostly flat", 9: "Flat with one prominent hill", 10: "Gentle hills", 11: "A single mountain", 12: "Mountainous and rugged"}

// 2d6
var islandVegetation = map[int]string{2: "One single palm tree", 3: "Barren: rocks", 4: "Swampy grasslands", 5: "Thorny scrub", 6: "Sand and meadows", 7: "Fertile grasslands", 8: "Fertile grasslands", 9: "Fields and woods", 10: "Small pine forests", 11: "Forests and meadows", 12: "Rainforest"}

// 2d6
var islandProminentNaturalFeature = map[int]string{2: "Volcanic", 3: "Huge sand dunes", 4: "Numerous sea stacks", 5: "Black sand beaches", 6: "Lagoon", 7: "Azure water, white sand", 8: "Waterfall(s)", 9: "Fruit trees", 10: "Mangrove forest", 11: "Sea cliffs", 12: "A huge skull-shaped rock"}

// 2d6
var islandNotableAnimalInhabitants = map[int]string{2: "Sharks. So many sharks.", 3: "Glowing jellyfish swarms", 4: "Monkeys", 5: "Fist sized orange beetles", 6: "Sleek black seals", 7: "Nothing to write home about", 8: "Pigs and boars (large)", 9: "Dog-sized lizards", 10: "Large seabird colony", 11: "Large flightless birds", 12: "Dancing jewel-colored birds"}

// 2d6
var islandInhabitants = map[int]string{2: "A witch", 3: "Pearl divers", 4: "A robust but primitive city", 5: "Small primitive village", 6: "Recently struck camps", 7: "Uninhabited", 8: "Don't know, but they're all dead", 9: "Religious adherents", 10: "Farming settlers", 11: "Military", 12: "A hidden pirate town"}

// 2d6
var islandBestPlaceToHideTreasure = map[int]string{2: "Old ruins", 3: "A large rock shaped like a ship", 4: "Under a huge lightning scarred oak", 5: "A large natural rock arch", 6: "An old lava tunnel", 7: "Somewhere else", 8: "Two palm trees cross in an X", 9: "A large rock resembling a monkey", 10: "A sink hole", 11: "Near the strange ancient statues", 12: "In the sea caves"}

// 2d6
var islandAnythingWorthAFortune = map[int]string{2: "Wealthy shipwreck survivors", 3: "Rare poisons", 4: "Oysters here make red pearls", 5: "Buried treasure... somewhere", 6: "An excellent natural harbor. Build a town!", 7: "Not really. Good fishing though.", 8: "Not exactly, but there's a rich trade routenearby", 9: "Caves of bat guano", 10: "Wild horses", 11: "Whale spawning grounds", 12: "A naturally hidden, excellent harbor"}

// 2d6
var islandNaturalHazards = map[int]string{2: "Manchineel trees. Everywhere.", 3: "Large whirlpools frequently form", 4: "Quicksand", 5: "Low lying areas flood. Frequently.", 6: "Coral reefs", 7: "None", 8: "Dangerous rocky shallows", 9: "Dangerous currents", 10: "Frequent earthquakes", 11: "Animals infect humans with hemorrhagic fever", 12: "A new volcano is about to erupt"}

// 2d6
var islandCurrentConflicts = map[int]string{2: "10 pirate ships here for a secret meeting", 3: "Imperial fleet firing upon inhabitants", 4: "Imperial military building an outpost", 5: "Pirates just raided, or are about to", 6: "Explorers hunting animals to extinction", 7: "None", 8: "Recent shipwreck. Survivors stranded.", 9: "Current inhabitants diseased", 10: "Religious order building a mission", 11: "A party of 6 ruthless treasure hunters is here (not including you...)", 12: "Rebels"}

// 2d6
var islandDarkness = map[int]string{2: "A kraken dwells here", 3: "Cultists", 4: "An ancient graveyard full of ghosts", 5: "A ghost ship haunts its waters", 6: "Ancient rock altar of sacrifice. Well used.", 7: "Nothing to speak of", 8: "Unspeakable monstrosities wash up on shore often", 9: "Animals that die here don't stay dead", 10: "Zombies walk at night", 11: "Blood spilled here animates", 12: "If you eat the food or drink the water here you can never leave."}

// -- ISLAND

// PIRATE --

// d12
var pirateTableA = []string{"rugged individual", "shady vagabond", "scurvy gutter worm", "famous captain", "escaped prisoner", "wanted cutthroat", "notorious privateer", "cloaked figure", "rum-drunk goon", "horrific wretch", "cocky explorer", "dread pirate lord"}

// d12
var pirateTableB = []string{"horrible breath", "unmatched beauty", "a hook for a hand", "a fancy hat", "a makeshift pegleg", "barnacle covered skin", "calico clothing", "filthy rags", "weather-worn doublet", "natty hair", "a pet bird", "bloody bandages"}

// d12
var pirateTableC = []string{"pair of fists [d2]", "broken bottle [d2]", "bullwhip [d4]", "hook [d4]", "rusty knife [d4]", "evil cutlass [d6]", "ornate scimitar [d6]", "officer's rapier [d8]", "ancient crossbow [d8]", "flintlock pistol [2d4]", "musket [2d6]", "blunderbuss [d4/d10]"}

// d36
var pirateMaleName = []string{"Esteban", "Richard", "Hendrik", "Raymond", "John", "Edmund", "Charles", "Peter", "Olivier", "Barth", "Henry", "Roger", "Don", "Martín", "Louis", "Fredrick", "Willem", "Nicholas", "Jerry", "Edward", "Alvaro", "Gaspar", "Francisco", "Johan", "Carlos", "Francis", "Jacques", "Jack", "François", "Silas", "Thomas", "Jacob", "Juan", "Philippe", "Jean", "William/Billy"}

// d36
var pirateFemaleName = []string{"Bridget", "Juliette", "Esther", "Rose", "Beatrice", "Olive", "Antonia", "Charlotte", "Isabel", "Adine", "Angela", "Cécile", "Edwidge", "Catalina", "Elizabeth", "Madeleine", "Anastasia", "Emma", "Mary", "Francisca", "Ana", "Agnes", "Marie", "Eleanor", "Anne", "Henrietta", "Alice", "Margaret", "Jeannette", "Camela", "Catherine", "Ursula", "Anette", "Gabriel", "Esme", "Marion"}

// d36
var pirateSurname = []string{"Pérez", "Thompson", "Jansen", "Williams", "Alva", "Dubois", "Leon", "Brown", "Jones", "Johnson", "Thatch", "Davies", "Archer", "Blanc", "Evans", "Wright", "Smith", "Wilson", "Bernard", "Roberts", "White", "Jean", "Santiago", "Morel", "Rodríguez", "Garcia", "Robinson", "López", "Baker", "Black", "Bonnet", "Walker", "Martin", "Jackson", "Diaz", "Taylor"}

// d36
var pirateNickname = []string{"Sir/Madam", "Sea", "Turtle", "Siren", "Red", "One-Eye/Arm/Leg", "Crimson", "Blue", "Water", "Skull", "Tall-Tale", "Old", "Blood", "Mr./Mrs./Miss", "Gunpowder", "King/Queen", "Bow-legged", "Fish", "Whale", "Bones", "Squid", "Scurvy", "Bilge", "Shark", "Heart", "The Bride/Groom", "Black", "White", "Death", "Dark", "Devil", "Knife", "Claw", "Rat", "Green", "Planktooth"}

// -- PIRATE

// JOBS & QUESTS --

// d20
var jobsHooks = []string{"Capture a prize.", "Raid a port.", "Treasure hunt.", "Find someone/something.", "Capture, interrogate, or kill someone.", "Establish a new safe house, base, or hideout.", "Sink a ship.", "Steal an item.", "Spy on someone.", "Scavenge a wreck.", "Explore or defend a location.", "Sneak into somewhere.", "Steal something.", "Meet someone.", "Establish trading or fencing relations with someone.", "Free or rescue a person or group from a location.", "Sabotage something.", "Defend a location.", "Free captives or prisoners.", "Escape."}

// d10
var jobsKeyNPC = []string{"A government leader.", "A famous pirate captain.", "The governor's daughter or son.", "The local drunk.", "A shady figure.", "Some cultists.", "An old friend.", "Important looking pirates.", "A love interest.", "A mermaid."}

// d6
var jobsRequirements = []string{"Stay out of combat.", "Kill only one person.", "Kill everyone.", "Make it look like someone else did it.", "Bring a hostage back alive.", "Vital information must be obtained."}

// d6
var jobsTimeRestraint = []string{"Immediately!", "Within the hour.", "By tonight.", "By tomorrow.", "Next week.", "This month."}

// d6
var jobsComplication = []string{"It's a setup.", "The intel is wrong.", "Another crew is already on the job.", "Some who's not supposed to be there is.", "The reward is suspiciously big.", "Everything is proceeding as planned."}

// -- JOBS & QUESTS

// d20
var rumors = []string{"A nation's navy is coming in 3d6 days. Everyone is on edge.", "A rich group of merchants are hiring mercenaries to escort them.", "A treasure fleet is rumored to be departing soon.", "An important NPC is recruiting people to explore or scout an area.", "A group of leaders is meeting soon. The PCs have been summoned.", "Someone the PCs know has been killed or captured.", "2d12 pirates are in town spending the earnings from their latest prize.", "Pirates, marines, or revolutionaries are planning to raid a port.", "There are rumors of a rich prize nearby. Crews are recruiting.", "Someone uncovered a clue or map to a fabled treasure and need brains, muscle, or both.", "A corpse has been found in an alley or on the beach. Everyone is suspicious.", "A ship with valuable cargo wrecked nearby.", "A lonely widow (20+d50 years old) is looking for a date to the Governor's Ball.", "It is not safe here at night. The reasons given sound like tall tales or ghost stories.", "A recent fire destroyed a ship, a port, or a farmstead.", "Two allied factions are now at war.", "There is a party tonight to honor a deceased NPC.", "The governor's adult daughter/son is the most beautiful person ever seen.", "A group of cultists have been up to some unusual activity recently.", "Undead have destroyed yet another settlement."}

// 2d6
var reaction = map[int]string{2: "bloodthirsty", 3: "bloodthirsty", 4: "angered or alarmed", 5: "angered or alarmed", 6: "angered or alarmed", 7: "indifferent, neutral", 8: "indifferent, neutral", 9: "almost friendly", 10: "almost friendly", 11: "helpful", 12: "helpful"}
