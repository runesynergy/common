// AUTOMATICALLY GENERATED - DO NOT MODIFY
//go:generate go fmt .

package enum

import (
	"common.runesynergy.dev/enum/combat"
	"common.runesynergy.dev/enum/equipment"
	"common.runesynergy.dev/enum/stat"
)

// Aliases

type (
	Sound             = string
	Model             = string
	Options           = map[int]string
	SkillRequirements = map[stat.Skill]stat.Level
	WeaponAnimations  = map[combat.Action]*Animation
	WeaponSounds      = map[combat.Action]Sound
)

// External asset references

type Animation struct {
	Ref string
}

func (a *Animation) String() string {
	return a.Ref
}

var _animations []*Animation
var _animationByRef = make(map[string]*Animation)

func NewAnimation(v *Animation, extras ...any) (animation *Animation) {
	ref := v.Ref
	if ref == "" {
		panic("missing ref")
	}
	animation = _animationByRef[ref]
	if animation == nil {
		_animationByRef[ref] = v
		animation = v
	} else {
		*animation = *v
	}
	return
}

func AnimationByRef(ref string) (animation *Animation) {
	return _animationByRef[ref]
}

func Animations() (animations []*Animation) {
	return _animations
}

// Data types

type NPC struct {
	Ref      string
	Name     string
	Examine  string
	Category []string
}

func (n *NPC) String() string {
	return n.Ref
}

var _npcs []*NPC
var _npcByRef = make(map[string]*NPC)

func NewNPC(v *NPC, extras ...any) (npc *NPC) {
	ref := v.Ref
	if ref == "" {
		panic("missing ref")
	}
	npc = _npcByRef[ref]
	if npc == nil {
		_npcByRef[ref] = v
		npc = v
	} else {
		*npc = *v
	}
	return
}

func NPCByRef(ref string) (npc *NPC) {
	return _npcByRef[ref]
}

func NPCs() (npcs []*NPC) {
	return _npcs
}

type Object struct {
	Ref  string
	Name string
}

func (o *Object) String() string {
	return o.Ref
}

var _objects []*Object
var _objectByRef = make(map[string]*Object)

func NewObject(v *Object, extras ...any) (object *Object) {
	ref := v.Ref
	if ref == "" {
		panic("missing ref")
	}
	object = _objectByRef[ref]
	if object == nil {
		_objectByRef[ref] = v
		object = v
	} else {
		*object = *v
	}
	return
}

func ObjectByRef(ref string) (object *Object) {
	return _objectByRef[ref]
}

func Objects() (objects []*Object) {
	return _objects
}

type itemFlag int

const (
	itemFlagCertificate itemFlag = 1 << iota
)

type ItemIcon struct {
	Zoom  int
	Pitch int
	Yaw   int
	Roll  int
	OffX  int
	OffY  int
}

type ItemOverride struct {
	Count int
	Use   *Item
}

type Item struct {
	Ref              string
	Name             string
	Examine          string
	Stackable        bool
	Members          bool
	Tradeable        bool
	QuestRelated     bool
	Cost             uint32
	Weight           int32
	Team             uint8
	InventoryOptions Options
	Options          Options
	Icon             ItemIcon
	Model            Model
	Recolors         []ModelRecolor
	Overrides        []ItemOverride

	linked *Item
	flags  itemFlag

	axe       *Axe
	weapon    *Weapon
	equipment *Equipment
}

func (i *Item) String() string {
	return i.Ref
}

var _items []*Item
var _itemByRef = make(map[string]*Item)

func NewItem(v *Item, extras ...any) (item *Item) {
	ref := v.Ref
	if ref == "" {
		panic("missing ref")
	}
	item = _itemByRef[ref]
	if item == nil {
		_itemByRef[ref] = v
		item = v
	} else {
		*item = *v
	}
	for _, extra := range extras {
		switch extra := extra.(type) {
		case *Axe:
			extra.Item = item
			item.axe = extra
		case *Weapon:
			extra.Item = item
			item.weapon = extra
		case *Equipment:
			item.equipment = extra
		}
	}
	if item.Tradeable && !item.Stackable {
		item.linked = NewItemCertificate(item)
	}
	return
}

func (i *Item) Axe() (axe *Axe, ok bool) {
	axe, ok = i.axe, axe != nil
	return
}

func (i *Item) Weapon() (weapon *Weapon, ok bool) {
	weapon, ok = i.weapon, weapon != nil
	return
}

func (i *Item) Equipment() (equipment *Equipment, ok bool) {
	equipment, ok = i.equipment, equipment != nil
	return
}

func ItemByRef(ref string) (item *Item) {
	return _itemByRef[ref]
}

func Items() (items []*Item) {
	return _items
}

type Axe struct {
	*Item
	Requirement int
	Bonus       int
}

type Weapon struct {
	*Item
	Type       *WeaponType
	Speed      int
	Animations WeaponAnimations
	Sounds     WeaponSounds
}

type Equipment struct {
	Slot              equipment.Slot
	Bonuses           Bonuses
	SkillRequirements SkillRequirements
	WomanModels       []Model
	WomanHeadModels   []Model
	WomanModelOffset  int
	ManModels         []Model
	ManHeadModels     []Model
	ManModelOffset    int
}

func (i *Item) HasCertificate() bool {
	return !i.IsCertificate() && i.linked != nil
}

func (i *Item) IsCertificate() bool {
	return (i.flags & itemFlagCertificate) != 0
}

func (i *Item) Certificate() (cert *Item) {
	if i.HasCertificate() {
		cert = i.linked
	}
	return
}

func (i *Item) Uncertificate() (item *Item) {
	if i.IsCertificate() {
		item = i.linked
	}
	return
}

func NewItemCertificate(obj *Item) (note *Item) {
	note = NewItem(&Item{
		Ref:       "cert_" + obj.Ref,
		Name:      obj.Name,
		Examine:   obj.Examine,
		Tradeable: true,
		Stackable: true,
		Cost:      obj.Cost,
		linked:    obj,
		flags:     itemFlagCertificate,
	})
	return
}

type WeaponType struct {
	Ref        string
	Animations WeaponAnimations
	Sounds     WeaponSounds
	Options    []combat.Option
}

func (w *WeaponType) String() string {
	return w.Ref
}

var _weapontypes []*WeaponType
var _weapontypeByRef = make(map[string]*WeaponType)

func NewWeaponType(v *WeaponType, extras ...any) (weapontype *WeaponType) {
	ref := v.Ref
	if ref == "" {
		panic("missing ref")
	}
	weapontype = _weapontypeByRef[ref]
	if weapontype == nil {
		_weapontypeByRef[ref] = v
		weapontype = v
	} else {
		*weapontype = *v
	}
	return
}

func WeaponTypeByRef(ref string) (weapontype *WeaponType) {
	return _weapontypeByRef[ref]
}

func WeaponTypes() (weapontypes []*WeaponType) {
	return _weapontypes
}

// Composite types

type Tree struct {
	*Object
	Requirement int
	Difficulty  int
}

type ModelRecolor struct {
	Src, Dst uint16
}

type Bonuses struct {
	AttackStab     int16
	AttackSlash    int16
	AttackCrush    int16
	AttackMagic    int16
	AttackRanged   int16
	DefenceStab    int16
	DefenceSlash   int16
	DefenceCrush   int16
	DefenceMagic   int16
	DefenceRanged  int16
	MeleeStrength  int16
	RangedStrength int16
	MagicStrength  int16
	Prayer         int16
}
