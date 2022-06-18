import { Soldier } from './soldier'

export class Player {
    private id: string = ''
    private nickName: string = ''

    private soldiers: Soldier[] = []
    // 是否到玩家回合
    private isRound: boolean = false
}
