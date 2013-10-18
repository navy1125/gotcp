package Cmd

import ()

/// 时间指令
const TIME_USERCMD = 2
const GM_USERCMD = 62

type stTimerUserCmd struct {
	StNullUserCmd
}

/// 网关向用户发送游戏时间
const GAMETIME_TIMER_USERCMD_PARA = 1

type StGameTimeTimerUserCmd struct {
	stTimerUserCmd
	QwGameTime uint64 /**< 游戏时间 */
	DwTempID   uint64 /**< 断线重连随机数 */
}

func NewStGameTimeTimerUserCmd() *StGameTimeTimerUserCmd {
	cmd := &StGameTimeTimerUserCmd{}
	cmd.ByCmd = TIME_USERCMD
	cmd.ByParam = GAMETIME_TIMER_USERCMD_PARA
	return cmd
}

/// 网关向用户发送游戏时间
const REQUESTUSERGAMETIME_TIMER_USERCMD_PARA = 2

type StRequestUserGameTimeTimerUserCmd struct {
	stTimerUserCmd
}

func NewStRequestUserGameTimeTimerUserCmd() *StRequestUserGameTimeTimerUserCmd {
	cmd := &StRequestUserGameTimeTimerUserCmd{}
	cmd.ByCmd = TIME_USERCMD
	cmd.ByParam = REQUESTUSERGAMETIME_TIMER_USERCMD_PARA
	return cmd
}

/// 网关向用户发送游戏时间
const USERGAMETIME_TIMER_USERCMD_PARA = 3

type StUserGameTimeTimerUserCmd struct {
	stTimerUserCmd
	qwGameTime uint64 /**< 用户游戏时间 */
	mac        uint64 /**< 用户MAC */
}

func NewStUserGameTimeTimerUserCmd() *StUserGameTimeTimerUserCmd {
	cmd := &StStUserGameTimeTimerUserCmd{}
	cmd.ByCmd = TIME_USERCMD
	cmd.ByParam = USERGAMETIME_TIMER_USERCMD_PARA
	return cmd
}
