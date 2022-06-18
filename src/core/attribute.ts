export interface AttributeOption {
    // 血量
    hp: number
    // 物理防御 physical defence
    pd: number
    // 法术防御 magical defence
    md: number
    // 物理攻击力 physical damage point
    pdp: number
    // 法术攻击力 magical damage point
    mdp: number
    // 攻击距离
    d: number
}

// 属性
export class Attribute {
    // 血量
    private hp: number
    // 物理防御 physical defence
    private pd: number
    // 法术防御 magical defence
    private md: number
    // 物理攻击力 physical damage point
    private pdp: number
    // 法术攻击力 magical damage point
    private mdp: number
    // 攻击距离
    private d: number

    constructor(opt: AttributeOption) {
        this.hp = opt.hp
        this.pd = opt.pd
        this.md = opt.md
        this.pdp = opt.pdp
        this.mdp = opt.mdp
        this.d = opt.d
    }
}
