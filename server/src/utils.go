package main

type IDType uint32
type MSTimeType = float64

const TICKS_PER_SECOND = 60
const SECONDS_PER_TICK = 1.0 / TICKS_PER_SECOND
const MS_PER_TICK float64 = SECONDS_PER_TICK * 1000.0

const MAX_TICKS_PER_UPDATE = 5