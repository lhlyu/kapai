export interface IRole {}

// 角色
export class Role {
    // 血量
    public hp: number = 0
    // 物理防御 physical defence
    public pd: number = 0
    // 法术防御 magical defence
    public md: number = 0
    // 物理攻击力 physical damage point
    public pdp: number = 0
    // 法术攻击力 magical damage point
    public mdp: number = 0
    // 攻击距离
    public d: number = 1
}

// 法师
export class Mage extends Role implements IRole {
    constructor() {
        super()
        this.hp = 10
        this.pd = 1
        this.md = 1
        this.pdp = 0
        this.mdp = 4
        this.d = 1
    }
}

// 战士
export class Warrior extends Role implements IRole {
    constructor() {
        super()
        this.hp = 16
        this.pd = 3
        this.md = 1
        this.pdp = 3
        this.mdp = 0
        this.d = 1
    }
}

// 牧师
export class Priest extends Role implements IRole {
    constructor() {
        super()
        this.hp = 11
        this.pd = 1
        this.md = 0
        this.pdp = 1
        this.mdp = 2
        this.d = 1
    }
}

// 射手
export class Shooter extends Role implements IRole {
    constructor() {
        super()
        this.hp = 12
        this.pd = 1
        this.md = 1
        this.pdp = 3
        this.mdp = 0
        this.d = 2
    }
}

// 刺客
export class Assassin extends Role implements IRole {
    constructor() {
        super()
        this.hp = 8
        this.pd = 0
        this.md = 0
        this.pdp = 5
        this.mdp = 0
        this.d = 1
    }
}

// 坦克
export class Tanker extends Role implements IRole {
    constructor() {
        super()
        this.hp = 20
        this.pd = 3
        this.md = 3
        this.pdp = 3
        this.mdp = 0
        this.d = 1
    }
}
