package pirateborg

import (
	"fmt"

	"github.com/genrpg/utils"
)

type HauntedSoul struct {
	Name            string    `json:"name"`
	Descrptn        string    `json:"description"`
	Features        []Feature `json:"features"`
	AdditionalClass PlayerClass
}

func NewHauntedSoul() PlayerClass {
	hauntedSoul := &HauntedSoul{
		AdditionalClass: GetClass(),
	}
	hauntedSoul.GetAilment()
	return hauntedSoul
}

func (hs *HauntedSoul) GetAilment() {
	r := utils.D(6)
	switch r {
	case 1:
		hs.GetGhost()
	case 2:
		hs.GetConduit()
	case 3:
		hs.GetEldritchMind()
	case 4:
		hs.GetZombie()
	case 5:
		hs.GetVampirism()
	case 6:
		hs.GetSkeleton()
	}
}

func (hs *HauntedSoul) GetDevilsLuck() string {
	return hs.AdditionalClass.GetDevilsLuck()
}

func (hs *HauntedSoul) GetClothingDie() int {
	return hs.AdditionalClass.GetClothingDie()
}
func (hs *HauntedSoul) GetHatDie() int {
	return hs.AdditionalClass.GetHatDie()
}
func (hs *HauntedSoul) GetWeaponDie() int {
	return hs.AdditionalClass.GetWeaponDie()
}

func (hs *HauntedSoul) GetFeatures() []Feature {
	return append(hs.Features, hs.AdditionalClass.GetFeatures()...)
}

func (hs *HauntedSoul) GetHPDie() int {
	return hs.AdditionalClass.GetHPDie()
}

func (hs *HauntedSoul) GetStatMods() Stats {
	return hs.AdditionalClass.GetStatMods()
}

func (hs *HauntedSoul) String() string {
	return fmt.Sprintf("%s | %s", hs.Name, hs.AdditionalClass.String())
}

func (hs *HauntedSoul) StartingFeatureBlurb() string {
	return ""
}

func (hs *HauntedSoul) Level() int {
	return hs.AdditionalClass.Level()
}

func (hs *HauntedSoul) LevelUp() {
	hs.AdditionalClass.LevelUp()
}

func (hs *HauntedSoul) GetItems() []Item {
	return hs.AdditionalClass.GetItems()
}

func (hs *HauntedSoul) GetWeapons() []Weapon {
	return hs.AdditionalClass.GetWeapons()
}

func (hs *HauntedSoul) Description() string {
	return fmt.Sprintf("%s\n%s", hs.Descrptn, hs.AdditionalClass.Description())
}

type HSFeature struct {
	Name        string
	Description string
}

func (hf HSFeature) String() string {
	return fmt.Sprintf("%s | %s", hf.Name, hf.Description)
}

func (hs *HauntedSoul) GetGhost() {
	hs.Name = "Ghost"
	hs.Descrptn = "You are a lost soul from beyond the grave inhabiting the body of another."
	hs.Features = []Feature{
		HSFeature{Name: "Terrify", Description: "Once per night, you can apparate and terrify a target with your ghostly visage: test SPIRIT DR14 to deal " + B("d12") + " damage."},
		HSFeature{Name: "Possession", Description: "If you are reduced to 1 HP or are somehow exorcised from your vessel, you become incorporeal and must find a new host in " + B("d12") + " hours or dissipate into the void. Possessing a new target: test SPIRIT DR16 (limit one attempt per target)."},
	}
}

func (hs *HauntedSoul) GetConduit() {
	hs.Name = "Conduit"
	hs.Descrptn = "Restless spirits often use you to communicate with the corporeal world."
	hs.Features = []Feature{HSFeature{Name: "Arcane Ritual", Description: "Every day at dawn, roll for a random " + B("Arcane Ritual (pg. 64)") + ". You can use that ritual once without testing SPIRIT, after which the spirits leave your body. If you do not cast the Ritual by dawn you lose " + B("d2") + " HP."}}
}

func (hs *HauntedSoul) GetEldritchMind() {
	hs.Name = "Eldritch Mind"
	hs.Descrptn = "Something dark and terrible wells underneath your conscious thoughts. You are drawn to The Deep and are prone to nightmarish visions of antediluvian horrors."
	hs.Features = []Feature{HSFeature{Name: "Evil Thoughts", Description: "Every time combat starts, test PRESENCE DR12 to attempt to control your evil thoughts. Fail: They are coming! You panic. Your attacks are +4 DR to hit, but you may make 2 a turn. Pass: You harness the fear. Attacks are -2 DR to hit for this combat."}}
}

func (hs *HauntedSoul) GetZombie() {
	hs.Name = "Zombie"
	hs.Descrptn = "You are secretly a zombie."
	hs.Features = []Feature{HSFeature{Name: "ZOMBIE!!!", Description: "Your flesh slowly rots and falls from your skin. You lose " + B("d2") + " HP every day that you don't consume human flesh. Eating a human brain restores " + B("2d6") + " HP. As long as your brain is intact, you maintain control of all of your body parts, attached or not."}}
}

func (hs *HauntedSoul) GetVampirism() {
	hs.Name = "Vampire"
	hs.Descrptn = "You have been infected by a fiendish creature."
	hs.Features = []Feature{
		HSFeature{Name: "Thirst", Description: "You cannot be healed by resting or other normal means, but drinking fresh animal blood restores " + B("d2") + " HP and fresh human blood restores " + B("d6") + " HP. You can only recover HP from drinking blood once an hour."},
		HSFeature{Name: "Night Creature", Description: "You are faster at night (move 40' a turn), but feel hungover in sunlight (-2 to TOUGHNESS)."},
	}
}

func (hs *HauntedSoul) GetSkeleton() {
	hs.Name = "Skeleton"
	hs.Descrptn = "You are an undead skeleton reanimated by unknowable dark magic, but somehow possessing free will."
	hs.Features = []Feature{
		HSFeature{Name: "Bone Dance", Description: "If you are killed, reroll the damage dice that killed you. On a 3 or less, your bones reform and you return to life with 1 HP."},
	}
}
