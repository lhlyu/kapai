import type { IRole } from './role'
import { Warrior } from './role'

export enum Position {
    TOP = 1,
    CENTER = 2,
    BOTTOM = 1
}

// 士兵
export class Soldier {
    // 位置
    private position: Position
    // 角色
    private role: IRole

    constructor(p: Position = Position.CENTER, role: IRole = Warrior) {
        this.position = p
        this.role = role
    }
}
