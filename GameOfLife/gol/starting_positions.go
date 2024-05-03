package gameoflife

/*
Options for starting state of cellular automata,
define as if top left of structure is x,y=0,0
this will be shifted to the center
*/

// STILL LIFES
var BLOCK = [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}

// OSCILLATORS
var BLINKER = [][]int{{0, 0}, {0, 1}, {0, 2}}

// SPACESHIPS
var GLIDER = [][]int{{0, 1}, {1, 2}, {2, 2}, {2, 1}, {2, 0}}
