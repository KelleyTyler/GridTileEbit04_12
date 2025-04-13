package userinterface

/*
coming up with a UI interface that is a generic object;
*/
type UI_Object interface {
	init() //initialize void;
	draw()
	update()
	getType() string //might want to change this output to like an int or something using a golang equivelant of an enum;
}
