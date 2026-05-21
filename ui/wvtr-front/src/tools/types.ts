
enum AffixType {
    Life = 0,
}

enum HeroTakeDamageStatus {
    TookDamage = 1 << 0,
    Dodged = 1 << 1,
    Blocked = 1 << 2,
    Died = 1 << 3,
    Crit = 1 << 4,
}

enum TargetType {
    Self = 1,
    Enemy,
    Friends,
}

enum SkillType {
    Unique = 0,
    Active,
}

enum SkillID {
    Lucky = 0,
    GoodRest,
    SecondWind,
    Prodigy,
    Berserk,
    Trickster,
    FastLearner,
    ElementalCursed,
    PhysicalCursed,
}

enum EncounterState {
    Home = 1,
    Travel,
    Fight,
    Neutral,
    Report,
    Error,
}

enum CurrencyType {
    Gold = 0,
    MScrap,
    LSCrap,
    CScrap,
}

enum EquipmentType {
    WeaponType = 0,
    ArmorType,
    OmamoriType,
}

enum InventoryToDo {
    Equip = 0,
    Spend,
}

class Damage {
    slashDmg: number = 0;
    bluntDmg: number = 0;
    pierceDmg: number = 0;
    fireDmg: number = 0;
    frostDmg: number = 0;
    lightningDmg: number = 0;
}

class StatsRange {
    min: number = 0
    max: number = 0
    value: number = 0
}

class Affix {
    name: string = ""
    ranges: StatsRange[] = []
    type: AffixType = 0
}

class Storable {
    name: string = ""
    iconURL: string = ""
}

class Currency extends Storable {
    type: CurrencyType | undefined
    // name: string,
    // iconURL: string,
}

class CurrencyOwned {
    numberOwned: number = 0
    currency: Currency | undefined
}

class Equipable extends Storable {
    realWeightScore: number = 0
    affixes: Affix[] = []
}

class Weapon extends Equipable {
    id: number = 0
    baseDamage: Damage | undefined
    baseCritRate: StatsRange | undefined
    baseAttackSpeed: StatsRange | undefined
}

class Armor extends Equipable {
    id: number = 0
    blockScore: StatsRange | undefined
    evadeScore: StatsRange | undefined
    baseResistancesRange: Damage | undefined
}

class Omamori extends Equipable {
    id: number = 0
}

class HeroEquipment {
    weapon: Weapon | undefined
    armor: Armor | undefined
    omamori: Omamori | undefined
}

class Inventory {
    weapons: Weapon[] = []
    armors: Armor[] = []
    omamoris: Omamori[] = []
    currencies: CurrencyOwned[] = []
}

class HeroAttributes {
    level: number = 0
    currentXP: number = 0
    xpBeforLvlUp: number = 0
    currentHP: number = 0

    //Attributes
    maxHP: number = 0
    strength: number = 0
    intelligence: number = 0
    dexterity: number = 0
    luck: number = 0

    //Growth rate
    hpgt: number = 0
    sgt: number = 0
    igt: number = 0
    dgt: number = 0
    lgt: number = 0

    //Defense
    blockScore: number = 0
    evadeScore: number = 0

    // Resistances
    blunt: number = 0
    pierce: number = 0
    slash: number = 0
    fire: number = 0
    frost: number = 0
    lighting: number = 0
}

class HeroClass {
    name: string = ""
    descritpion: string = ""
    class_icon_url: string = ""
}

class Reward {
    xp: number = 0
    loot: Inventory | undefined
}

class FieldActionDesc {
    fromH: Hero | undefined
    usedSKill: Skill | undefined
    targetH: Hero | undefined
    targetStatus: HeroTakeDamageStatus | undefined
    fromPVChange: number = 0
    targetPVChange: number = 0
}

class ExpeditionStepTimestamp {
    when: string = "" // time
    what: string = ""
    whatAction: FieldActionDesc | undefined
}

class ExpeditionStepResolveInfo {
    stepState: EncounterState | undefined
    timeline: ExpeditionStepTimestamp[] = []
    eTeam: Team | undefined
}

class ExpeditionDB {
    identifier: string = ""
    startedAt: string = ""
    whatHappened: ExpeditionStepResolveInfo[] = []
    expeditionRewards: Reward | undefined
}

class GameState {
    id: number = 0
    state: EncounterState | undefined
    currentExpedition: ExpeditionDB | undefined
}

class User {
    id: number = 0
    name: string = ""
    state: GameState | undefined
    inventory: Inventory | undefined
    currentTeam: Team | undefined
    lastActionTime: string = "" // time
    ownedHeroes: Hero[] = []
    discord_id: string = ""
}

class CurrentStepRequestMessage {
    id: number = 0
    time: number = 0
}

class Skill {
    identifier: SkillID | undefined
    name: string = ""
    skill_type: SkillType | undefined
    target_type: TargetType | undefined
    recuperation_duration: number = 0
    image_url: string = ""
    description: string = ""
}

class Hero {
    id: number = 0
    imageUrl: string = ""
    name: string = ""
    heroClass: HeroClass | undefined
    rank: string = ""
    attributes: HeroAttributes | undefined

    // skills
    weaponAttack: Skill | undefined
    uniqueSkill: Skill | undefined
    activeSkill: Skill | undefined

    // Items
    equipment: HeroEquipment | undefined

    // info that we save to request nanapi if we need to.
    id_w: string = ""
    id_al: number = 0
};

class Team {
    id: number = 0
    heroes: Hero[] = []
};

class Waifu {
    id: string = ""
    id_al: string = ""
    name_user_preferred: string = ""
    image_large: string = ""
    rank: string = ""
}

type ExpToLaunch = {
    cat: string,
    key: string,
    cost: string,
    costNumber: number,
}

type EquipmentToSpend = {
    eq_id: number,
    eq_type: EquipmentType,
}

export type {
    Hero,
    Team,
    GameState,
    User,
    ExpeditionStepResolveInfo,
    ExpeditionStepTimestamp,
    CurrentStepRequestMessage,
    ExpeditionDB,
    Waifu,
    FieldActionDesc,
    Inventory,
    Equipable,
    Weapon,
    Armor,
    Omamori,
    Storable,
    ExpToLaunch,
    EquipmentToSpend,
};

export {
    EncounterState,
    HeroTakeDamageStatus,
    EquipmentType,
    InventoryToDo,
}

