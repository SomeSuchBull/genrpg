package mothership

type SkillType string

const (
	SkillTypeTrained SkillType = "Trained"
	SkillTypeExpert  SkillType = "Expert"
	SkillTypeMaster  SkillType = "Master"
)

type Skill struct {
	Name            string
	Type            SkillType
	PrerequisiteFor []Skill
	Prerequisites   []Skill
}

func (s Skill) String() string {
	return s.Name
}

var (
	//
	MasterSkills = []Skill{}
	// Master Skills
	Sophontology           = Skill{Type: SkillTypeMaster, Name: "Sophontology"}
	ExoBiology             = Skill{Type: SkillTypeMaster, Name: "Exobiology"}
	Surgery                = Skill{Type: SkillTypeMaster, Name: "Surgery"}
	Planetology            = Skill{Type: SkillTypeMaster, Name: "Planetology"}
	Robotics               = Skill{Type: SkillTypeMaster, Name: "Robotics"}
	Engineering            = Skill{Type: SkillTypeMaster, Name: "Engineering"}
	Cybernetics            = Skill{Type: SkillTypeMaster, Name: "Cybernetics"}
	ArtificialIntelligence = Skill{Type: SkillTypeMaster, Name: "Artificial intelligence"}
	HyperSpace             = Skill{Type: SkillTypeMaster, Name: "Hyperspace"}
	Xenoesotericism        = Skill{Type: SkillTypeMaster, Name: "Xenoesotericism"}
	Command                = Skill{Type: SkillTypeMaster, Name: "Command"}
	//
	ExpertSkills = []Skill{}
	// Expert Skills
	Psychology         = Skill{Type: SkillTypeExpert, Name: "Psychology", PrerequisiteFor: []Skill{Sophontology}}
	Pathology          = Skill{Type: SkillTypeExpert, Name: "Pathology", PrerequisiteFor: []Skill{ExoBiology, Surgery}}
	FieldMedicine      = Skill{Type: SkillTypeExpert, Name: "Field medicine", PrerequisiteFor: []Skill{Surgery}}
	Ecology            = Skill{Type: SkillTypeExpert, Name: "Ecology", PrerequisiteFor: []Skill{Planetology}}
	AsteroidMining     = Skill{Type: SkillTypeExpert, Name: "Asteroid mining", PrerequisiteFor: []Skill{Planetology}}
	MechanicalRepair   = Skill{Type: SkillTypeExpert, Name: "Mechanical repair", PrerequisiteFor: []Skill{Robotics, Engineering, Cybernetics}}
	Explosives         = Skill{Type: SkillTypeExpert, Name: "Explosives", PrerequisiteFor: []Skill{}}
	Pharmacology       = Skill{Type: SkillTypeExpert, Name: "Pharmacology", PrerequisiteFor: []Skill{}}
	Hacking            = Skill{Type: SkillTypeExpert, Name: "Hacking", PrerequisiteFor: []Skill{ArtificialIntelligence}}
	Piloting           = Skill{Type: SkillTypeExpert, Name: "Piloting", PrerequisiteFor: []Skill{HyperSpace, Command}}
	Physics            = Skill{Type: SkillTypeExpert, Name: "Physics", PrerequisiteFor: []Skill{HyperSpace}}
	Mysticism          = Skill{Type: SkillTypeExpert, Name: "Mysticism", PrerequisiteFor: []Skill{HyperSpace, Xenoesotericism}}
	WildernessSurvival = Skill{Type: SkillTypeExpert, Name: "Wilderness survival", PrerequisiteFor: []Skill{}}
	FireArms           = Skill{Type: SkillTypeExpert, Name: "Firearms", PrerequisiteFor: []Skill{Command}}
	HandToHandCombat   = Skill{Type: SkillTypeExpert, Name: "Hand-to-hand combat", PrerequisiteFor: []Skill{}}
	//
	TrainedSkills = []Skill{}
	// Trained Skills
	Linguistics         = Skill{Type: SkillTypeTrained, Name: "Linguistics", PrerequisiteFor: []Skill{Psychology}}
	Zoology             = Skill{Type: SkillTypeTrained, Name: "Zoology", PrerequisiteFor: []Skill{Psychology, Pathology, FieldMedicine}}
	Botany              = Skill{Type: SkillTypeTrained, Name: "Botany", PrerequisiteFor: []Skill{Psychology, Pathology, Ecology, WildernessSurvival}}
	Geology             = Skill{Type: SkillTypeTrained, Name: "Geology", PrerequisiteFor: []Skill{AsteroidMining}}
	IndustrialEquipment = Skill{Type: SkillTypeTrained, Name: "Industrial equipment", PrerequisiteFor: []Skill{AsteroidMining, MechanicalRepair}}
	JuryRigging         = Skill{Type: SkillTypeTrained, Name: "Jury-rigging", PrerequisiteFor: []Skill{MechanicalRepair, Explosives}}
	Chemistry           = Skill{Type: SkillTypeTrained, Name: "Chemistry", PrerequisiteFor: []Skill{Explosives, Pharmacology}}
	Computers           = Skill{Type: SkillTypeTrained, Name: "Computers", PrerequisiteFor: []Skill{Hacking}}
	ZeroG               = Skill{Type: SkillTypeTrained, Name: "Zero-G", PrerequisiteFor: []Skill{Piloting}}
	Mathematics         = Skill{Type: SkillTypeTrained, Name: "Mathematics", PrerequisiteFor: []Skill{Physics}}
	Art                 = Skill{Type: SkillTypeTrained, Name: "Art", PrerequisiteFor: []Skill{Mysticism}}
	Archeology          = Skill{Type: SkillTypeTrained, Name: "Archeology", PrerequisiteFor: []Skill{Mysticism}}
	Theology            = Skill{Type: SkillTypeTrained, Name: "Theology", PrerequisiteFor: []Skill{Mysticism}}
	MilitaryTraining    = Skill{Type: SkillTypeTrained, Name: "Military training", PrerequisiteFor: []Skill{Explosives, WildernessSurvival, FireArms, HandToHandCombat}}
	Rimwise             = Skill{Type: SkillTypeTrained, Name: "Rimwise", PrerequisiteFor: []Skill{FireArms, HandToHandCombat}}
	Athletics           = Skill{Type: SkillTypeTrained, Name: "Athletics", PrerequisiteFor: []Skill{HandToHandCombat}}
)

func init() {
	// Set prerequisites for Master Skills
	Sophontology.Prerequisites = []Skill{Psychology}
	ExoBiology.Prerequisites = []Skill{Pathology}
	Surgery.Prerequisites = []Skill{Pathology, FieldMedicine}
	Planetology.Prerequisites = []Skill{Ecology, AsteroidMining}
	Robotics.Prerequisites = []Skill{MechanicalRepair}
	Engineering.Prerequisites = []Skill{MechanicalRepair}
	Cybernetics.Prerequisites = []Skill{MechanicalRepair}
	ArtificialIntelligence.Prerequisites = []Skill{Hacking}
	HyperSpace.Prerequisites = []Skill{Piloting, Physics, Mysticism}
	Xenoesotericism.Prerequisites = []Skill{Mysticism}
	Command.Prerequisites = []Skill{Piloting, FireArms}
	// Set prerequisites for Expert Skills
	Psychology.Prerequisites = []Skill{Linguistics, Zoology, Botany}
	Pathology.Prerequisites = []Skill{Zoology, Botany}
	FieldMedicine.Prerequisites = []Skill{Zoology, Botany}
	Ecology.Prerequisites = []Skill{Botany, Geology}
	AsteroidMining.Prerequisites = []Skill{Geology, IndustrialEquipment}
	MechanicalRepair.Prerequisites = []Skill{IndustrialEquipment, JuryRigging}
	Explosives.Prerequisites = []Skill{JuryRigging, Chemistry, MilitaryTraining}
	Pharmacology.Prerequisites = []Skill{Chemistry}
	Hacking.Prerequisites = []Skill{Computers}
	Piloting.Prerequisites = []Skill{ZeroG}
	Physics.Prerequisites = []Skill{Mathematics}
	Mysticism.Prerequisites = []Skill{Art, Archeology, Theology}
	WildernessSurvival.Prerequisites = []Skill{Botany, MilitaryTraining}
	FireArms.Prerequisites = []Skill{MilitaryTraining, Rimwise}
	HandToHandCombat.Prerequisites = []Skill{MilitaryTraining, Rimwise, Athletics}
	// Fill Skills slices
	MasterSkills = []Skill{Sophontology, ExoBiology, Surgery, Planetology, Robotics, Engineering, Cybernetics,
		ArtificialIntelligence, HyperSpace, Xenoesotericism, Command}
	ExpertSkills = []Skill{Psychology, Pathology, FieldMedicine, Ecology, AsteroidMining, MechanicalRepair, Explosives,
		Pharmacology, Hacking, Piloting, Physics, Mysticism, WildernessSurvival, FireArms, HandToHandCombat}
	TrainedSkills = []Skill{Linguistics, Zoology, Botany, Geology, IndustrialEquipment, JuryRigging, Chemistry,
		Computers, ZeroG, Mathematics, Art, Archeology, Theology, MilitaryTraining, Rimwise, Athletics}
}

type Equipment interface {
	String() string
}

type Item struct {
	Name        string
	Description string
	Cost        string
}

func (i Item) String() string {
	return i.Name
}

type Weapon struct {
	Name    string
	Cost    string
	Range   string
	Damage  string
	Shots   int
	Wound   string
	Special string
}

func (w Weapon) String() string {
	return w.Name
}

type Armor struct {
	Name        string
	Description string
	Cost        string
	ArmorPoints int
	O2          string
	Speed       string
	Special     string
}

func (a Armor) String() string {
	return a.Name
}

var (
	// Armor
	StandardCrewAttire  = Armor{Name: "Standard Crew Attire", Description: "Basic clothing.", Cost: "100cr", ArmorPoints: 1, O2: "None", Speed: "Normal"}
	Vaccsuit            = Armor{Name: "Vaccsuit", Description: "Designed for outer space operations", Cost: "10kcr", ArmorPoints: 3, O2: "12 hrs", Speed: "[-]", Special: "Includes short-range comms, headlamp, and radiation shielding. Decompression within 1d5 rounds if punctured."}
	HazardSuit          = Armor{Name: "Hazard Suit", Description: "Environmental protection while exploring unknown planets.", Cost: "4kcr", ArmorPoints: 5, O2: "1 hr", Speed: "Normal", Special: "Includes air filter, extreme heat/cold protection, hydration reclamation (1L of water lasts 4 days), short-range comms, headlamp, and radiation shielding."}
	StandardBattleDress = Armor{Name: "Standard Battle Dress", Description: "Lightly-plated armor worn by most marines.", Cost: "2kcr", ArmorPoints: 7, O2: "None", Speed: "Normal", Special: "Includes short-range comms."}
	AdvancedBattleDress = Armor{Name: "Advanced Battle Dress", Description: "Heavy armor for marines deployed in high combat offworld engagements.", Cost: "12kcr", ArmorPoints: 10, O2: "1 hr", Speed: "[-]", Special: "Includes short-range comms, body cam, headlamp, HUD, exoskeletal weave (Strength Checks [+]), and radiation shielding. Damage Reduction: 3."}
	// Weapons
	Ammo                     = Weapon{Name: "Ammo", Cost: "50cr", Special: "Per magazine/container"}
	BoardingAxe              = Weapon{Name: "Boarding Axe", Cost: "150cr", Range: "Adjacent", Damage: "2d10", Wound: "Gore[+]"}
	CombatShotgun            = Weapon{Name: "Combat Shotgun", Cost: "1.4kcr", Range: "Close", Damage: "4d10", Shots: 4, Wound: "Gunshot", Special: "1d10 DMG at Long Range or greater."}
	Crowbar                  = Weapon{Name: "Crowbar", Cost: "25cr", Range: "Adjacent", Damage: "1d5", Wound: "Blunt Force [+]", Special: "Grants [+] on Strength Checks to open jammed airlocks, lift heavy objects, etc."}
	Flamethrower             = Weapon{Name: "Flamethrower", Cost: "4kcr", Range: "Close", Damage: "2d10", Shots: 4, Wound: "Fire/Explosives[+]", Special: "Body Save [-] or be set on fire (2d10 DMG / round)."}
	FlareGun                 = Weapon{Name: "Flare Gun", Cost: "25cr", Range: "Long", Damage: "1d5", Shots: 2, Wound: "Fire/Explosives[-]", Special: "High intensity flare visible day and night from Long Range."}
	FoamGun                  = Weapon{Name: "Foam Gun", Cost: "500cr", Range: "Close", Damage: "1", Shots: 3, Wound: "Blunt Force", Special: "Body Save or become stuck. Strength Check [-] to escape."}
	FragGrenade              = Weapon{Name: "Frag Grenade", Cost: "400cr", Range: "Close", Damage: "3d10", Shots: 1, Wound: "Fire/Explosives", Special: "On a hit, damages all Adjacent to enemy."}
	GeneralPurposeMachineGun = Weapon{Name: "General-Purpose Machine Gun", Cost: "4.5kcr", Range: "Long", Damage: "4d10", Shots: 5, Wound: "Gunshot[+]", Special: "Two-handed. Heavy. Barrel can be maneuvered to fire around corners."}
	HandWelder               = Weapon{Name: "Hand Welder", Cost: "250cr", Range: "Adjacent", Damage: "1d10", Wound: "Bleeding", Special: "Can cut through airlock doors."}
	LaserCutter              = Weapon{Name: "Laser Cutter", Cost: "1.2kcr", Range: "Long", Damage: "1d100", Shots: 6, Wound: "Bleeding [+] or Gore [+]", Special: "Two-handed. Heavy. 1 round recharge between shots."}
	NailGun                  = Weapon{Name: "Nail Gun", Cost: "150cr", Range: "Close", Damage: "1d5", Shots: 32, Wound: "Bleeding"}
	PulseRifle               = Weapon{Name: "Pulse Rifle", Cost: "2.4kcr", Range: "Long", Damage: "3d10", Shots: 5, Wound: "Gunshot"}
	Revolver                 = Weapon{Name: "Revolver", Cost: "500cr", Range: "Close", Damage: "1d10+1", Shots: 6, Wound: "Gunshot"}
	RiggingGun               = Weapon{Name: "Rigging Gun", Cost: "350cr", Range: "Close", Damage: "1d10+2d10 when removed", Shots: 1, Wound: "Bleeding[+]", Special: "100m micro-filament. Body Save or become entangled."}
	Scalpel                  = Weapon{Name: "Scalpel", Cost: "50cr", Range: "Adjacent", Damage: "1d5", Wound: "Bleeding [+]"}
	SmartRifle               = Weapon{Name: "Smart Rifle", Cost: "5kcr", Range: "Extreme", Damage: "4d10 (AA)", Shots: 3, Wound: "Gunshot[+]", Special: "[-] on Combat Check when fired at Close Range."}
	SMG                      = Weapon{Name: "SMG", Cost: "1kcr", Range: "Long", Damage: "2d10", Shots: 5, Wound: "Gunshot", Special: "Can be fired one-handed."}
	StunBaton                = Weapon{Name: "Stun Baton", Cost: "150cr", Range: "Adjacent", Damage: "1d5", Wound: "Blunt Force", Special: "Body Save or stunned for 1 round."}
	TranqPistol              = Weapon{Name: "Tranq Pistol", Cost: "250cr", Range: "Close", Damage: "1d5", Shots: 6, Wound: "Blunt Force", Special: "If DMG dealt: enemy must Body Save or be unconscious 1d10 rounds."}
	Unarmed                  = Weapon{Name: "Unarmed", Range: "Adjacent", Damage: "Str/10", Wound: "Blunt Force"}
	Vibechete                = Weapon{Name: "Vibechete", Cost: "1kcr", Range: "Adjacent", Damage: "3d10 (AA)", Wound: "Bleeding + Gore", Special: "When dealing a Wound, roll on BOTH the Bleeding AND Gore columns."}
	// Items
	AssortedTools               = Item{Name: "Assorted Tools", Cost: "20cr", Description: "Wrenches, spanners, screwdrivers, etc. Can be used as weapons in a pinch (1d5 DMG)."}
	Automed                     = Item{Name: "Automed (x5)", Cost: "1.5kcr", Description: "Nanotech pills that assist your body in repairing Damage by granting Advantage to Body Saves meant to repel disease and poison, as well as attempts to heal from rest."}
	Battery                     = Item{Name: "Battery (High Power)", Cost: "500cr", Description: "Heavy duty battery used for powering laser cutters, salvage drones, and other items. Can be recharged in 1 hour if connected to power or in 6 hours with solar power. Add waterproofing (+500cr)."}
	Binoculars                  = Item{Name: "Binoculars", Cost: "150cr", Description: "20x magnification. Add night vision (+300cr) or thermal vision (+1kcr)."}
	Bioscanner                  = Item{Name: "Bioscanner", Cost: "3kcr", Description: "Long Range. Allows the user to scan for signs of life. Can tell the location of signs of life, but not what that life is. Blocked by some materials at the Warden’s discretion."}
	BodyCam                     = Item{Name: "Body Cam", Cost: "50cr", Description: "A camera worn on your clothing that can stream video back to a control center so your other crewmembers can see what you’re seeing. Add night vision (+300cr) or thermal vision (+1kcr)."}
	Chemlights                  = Item{Name: "Chemlights (x5)", Cost: "5cr", Description: "Small disposable glowsticks capable of dim illumination in a 1m radius"}
	CyberneticDiagnosticScanner = Item{Name: "Cybernetic Diagnostic Scanner", Cost: "2kcr", Description: "Allows the user to scan androids and other cybernetic organisms in order to diagnose any physical or mental issues they may be having. Often distrusted by androids."}
	ElectronicToolSet           = Item{Name: "Electronic Tool Set", Cost: "100cr", Description: "A full set of tools for doing detailed repair or construction work on electronics"}
	EmergencyBeacon             = Item{Name: "Emergency Beacon", Cost: "2kcr", Description: "A small device that sends up a flare and then emits a loud beep every few seconds. Additionally, sends out a call on all radio channels to ships or vehicles in the area, but can be blocked by a radio jammer."}
	Exoloader                   = Item{Name: "Exoloader", Cost: "100kcr", Description: "Open-air mechanical exoskeleton used for heavy lifting (up to 5000kg). Loader claws deal 1 Wound. User can only wear Standard Crew Attire or Standard Battle Dress while operating. Battery operated (48 hours of use)."}
	ExplosivesAndDetonator      = Item{Name: "Explosives & Detonator", Cost: "500cr", Description: "Explosive charge powerful enough to blow open an airlock. All organisms in Close Range must make a Body Save or take a Wound (Explosive). Detonator works at Long Range, but can be blocked by a radio jammer."}
	FirstAidKit                 = Item{Name: "First Aid Kit", Cost: "75cr", Description: "An assortment of dressings and treatments to help stop bleeding, bandage cuts, and treat other minor injuries."}
	Flashlight                  = Item{Name: "Flashlight", Cost: "30cr", Description: "Handheld or shoulder mounted. Illuminates 10m ahead of the user."}
	FoldableStretcher           = Item{Name: "Foldable Stretcher", Cost: "150cr", Description: "Portable stretcher that can fit within a rucksack. Allows the user to safely strap down the patient and carry them to a location where their wounds can be better treated. Unfolds to roughly 2m."}
	GeigerCounter               = Item{Name: "Geiger Counter", Cost: "20cr", Description: "Detects radiation and displays radiation levels."}
	HUD                         = Item{Name: "Heads-Up Display (HUD)", Cost: "100cr", Description: "Often worn by marines, the HUD allows the wearer to see through the body cams of others in their unit, and connect to any smart-link upgraded weapon."}
	InfraredGoggles             = Item{Name: "Infrared Goggles", Cost: "1.5kcr", Description: "Allows the wearer to see heat signatures, sometimes up to several hours old. Add night vision (+300cr)."}
	Jetpack                     = Item{Name: "Jetpack", Cost: "75kcr", Description: "Allows wearer to fly up to 100m high and up to a speed of 100km/hr for 2 hours on a tank of fuel. Deals 1d100[+] DMG if destroyed. Fuel can be refilled for 200cr."}
	LockpickSet                 = Item{Name: "Lockpick Set", Cost: "40cr", Description: "A highly advanced set of tools meant for hacking basic airlock and electronic door systems."}
	LongrangeComms              = Item{Name: "Long-range Comms", Cost: "1kcr", Description: "Rucksack-sized communication device for use in surface-to-ship communication."}
	MagBoots                    = Item{Name: "Mag-boots", Cost: "350cr", Description: "Grants a magnetic grip to the wearer, allowing them to easily walk on the exterior of a ship (in space, while docked, or free-floating), metal-based asteroids, or any other magnetic surface."}
	Medscanner                  = Item{Name: "Medscanner", Cost: "8kcr", Description: "Allows the user to scan a living or dead body to analyze it for disease or abnormalities, without having to do a biopsy or autopsy. Results may not be instantaneous and may require a lab for complete analysis."}
	MoHabUnit                   = Item{Name: "MoHab Unit", Cost: "1kcr", Description: "Tent, canteen, stove, rucksack, compass, and sleeping bag."}
	MRE                         = Item{Name: "MRE (x7)", Cost: "70cr", Description: "\"Meal, Ready-to-Eat.\" Self-contained, individual field rations in lightweight packaging. Each has sufficient sustenance for a single person for one day (does not include water)."}
	MylarBlanket                = Item{Name: "Mylar Blanket", Cost: "10cr", Description: "Lightweight blanket made of heat-reflective material. Often used for thermal regulation of patients suffering from extreme cold or other trauma."}
	OxygenTank                  = Item{Name: "Oxygen Tank", Cost: "50cr", Description: "When attached to a vaccsuit provides up to 12 hours of oxygen under normal circumstances, 4 hours under stressful circumstances. Explosive."}
	Paracord                    = Item{Name: "Paracord (50m)", Cost: "10cr", Description: "General purpose lightweight nylon rope."}
	PatchKit                    = Item{Name: "Patch Kit (x3)", Cost: "200cr", Description: "Repairs punctured and torn vaccsuits, restoring their space readiness. Patched vaccsuits have an AP of 1."}
	PersonalLocator             = Item{Name: "Personal Locator", Cost: "200cr", Description: "Allows crewmembers at a control center (or on the bridge of a ship) to track the location of the wearer."}
	PetOrganic                  = Item{Name: "Pet (Organic)", Cost: "200kcr", Description: "Small to medium-sized organic pet animal. Larger or rare pets cost 2d10x base pet cost."}
	PetSynthetic                = Item{Name: "Pet (Synthetic)", Cost: "15kcr", Description: "Small to medium-sized synthetic pet animal. Larger or rare pets cost 2d10x base pet cost."}
	PortableComputerTerminal    = Item{Name: "Portable Computer Terminal", Cost: "1.5kcr", Description: "Flat computer monitor, keyboard and interface which allows the user to hack into pre-existing computers and networks, as well as perform standard computer tasks."}
	RadiationPills              = Item{Name: "Radiation Pills (x5)", Cost: "200cr", Description: "Take 1d5 DMG and reduce your Radiation Level (see pg. 33.2) by 1 for 2d10 minutes."}
	RadioJammer                 = Item{Name: "Radio Jammer", Cost: "4kcr", Description: "Rucksack-sized device which, when activated, renders all radio signals within 100km incomprehensible."}
	Rebreather                  = Item{Name: "Rebreather", Cost: "500cr", Description: "When worn, filters toxic air and/or allows for underwater breathing for up to 20 minutes at a time without resurfacing. Can be connected to an oxygen tank."}
	Rucksack                    = Item{Name: "Rucksack", Cost: "50cr", Description: "Large, durable, waterproof backpack."}
	SalvageDrone                = Item{Name: "Salvage Drone", Cost: "10kcr", Description: "Battery operated remote controlled drone. Requires two hands to operate receiver. Can fly up to 450m high, to a distance of 3km from operator. Can run for 2 hours. Can record and transmit footage to receiver. If purchased separately, can be equipped with up to two of the following: binoculars, radio jammer, Geiger counter, laser cutter, medscanner, personal locator, infrared goggles, emergency beacon, cybernetic diagnostic scanner, bioscanner. Can carry up to 20-30kg."}
	SampleCollectionKit         = Item{Name: "Sample Collection Kit", Cost: "50cr", Description: "Used to research xenoflora and xenofauna in the field. Can take vital signs, DNA samples ,and collect other data on foreign material. Results may not be instantaneous and may require a lab for complete analysis."}
	ShortrangeComms             = Item{Name: "Short-range Comms", Cost: "100cr", Description: "Allows communication from ship-to-ship within a reasonable distance, as well as surface-to-surface within a dozen kilometers. Blocked by radio jammer."}
	SmartlinkAddOn              = Item{Name: "Smart-link Add-On", Cost: "10kcr", Description: "Grants remote viewing, recording, and operation of a ranged weapon as well as +5 DMG to the weapon."}
	Stimpak                     = Item{Name: "Stimpak", Cost: "1kcr ea.", Description: "Cures cryosickness, reduces Stress by 1, restores 1d10 Health, and grants [+] to all rolls for 1d10 min. Roll 1d10. If you roll under the amount of doses you’ve taken in the past 24 hours, make a Death Save."}
	WaterFiltrationDevice       = Item{Name: "Water Filtration Device", Cost: "50cr", Description: "Can pump 4 liters of filtered water per hour from even the most brackish swamps."}
)

// TODO: Trinkets and patches
