# Rambling thoughts of no particular order;

primarily about the way I'm implementing a GUI setup;


this is the currently existing/currently implemented "UI_Object" interface: I think it's kind of bloated and unhelpful;
```
type UI_Object interface {
	// Init0() //initialize void;
	Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error //--
	Init_Parents(Parent UI_Object) error                                                                              //--
	Draw(screen *ebiten.Image) error                                                                                  //--
	Redraw()                                                                                                          //--
	Update() error                                                                                                    //--
	Update_Unactive() error                                                                                           //

	Update_Any() (any, error) //
	Update_Ret_State_Redraw_Status() (uint8, bool, error)
	Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error)

	GetState() uint8                                                    //
	ToString() string                                                   //
	IsInit() bool                                                       //
	GetID() string                                                      //
	GetType() string                                                    //
	IsCursorInBounds() bool                                             //
	IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool //
	GetPosition_Int() (int, int)                                        //
	GetNumber_Children() int                                            //
	GetChild(index int) UI_Object
	AddChild(child UI_Object) error //
	RemoveChild(index int) error
	GetParent() UI_Object //
	HasParent() bool      //
	// getType() string //might want to change this output to like an int or something using a golang equivelant of an enum;
}
```
this is UI_Object_Style; a struct that helps me have some shorthand for all the variations in style I might want to have;
```
type UI_Object_Style struct {
	LabelColor       color.Color
	PanelColor       color.Color
	BorderColor      color.Color
	ButtonColor0     []color.Color
	ButtonColor1     []color.Color
	BorderThickness  float32
	TextColor        []color.Color
	Internal_Margins [4]uint8
	TextSizes        []int
	TextAlignMode    int
}

```
all these modules are connected through to the UI backend thing;
```
type UI_Backend struct {
	Settings                    *settings.GameSettings
	SoundSystem                 *gensound.Basic_SoundSystem
	Btn_Sounds                  [][]byte
	Textsrcs                    []*text.GoTextFaceSource
	Btn_Text_Mono, Btn_Text_Reg text.Face
	BtnColors0, BtnColors1      []color.Color
	Style                       UI_Object_Style
}
```


## Critique

This is a critique as I see it; or what I'm worried about having to do a lot of code wrangling/deleting files/git-rm-ing stuff;

---

### UI_Backend
```
type UI_Backend struct {
	Settings                    *settings.GameSettings
	SoundSystem                 *gensound.Basic_SoundSystem
	Btn_Sounds                  [][]byte
	Textsrcs                    []*text.GoTextFaceSource
	Btn_Text_Mono, Btn_Text_Reg text.Face
	BtnColors0, BtnColors1      []color.Color
	Style                       UI_Object_Style
}   
```
the problem I see with UI backend is complex; but basically it comes down to being unsure if it's a good way to implement sound/audio;
it feels like a bad way to go about that but it's nice in that I have a shorthand for hitting/activating sounds;
Text is another thing; perhaps I need to do some tests to see if there are any tradeoffs with having on-the-fly requests for a textface; but how to implement that as a call?
failing that just have an array and maybe a font file in an assets folder ready to go for some of it;
the main thing is just having the fonts properly sized when I need them to be;


sounds and textures are another set of 'valid' concerns; though I need to make sure that I'm using sounds effectively and not in a stupid way;
need to come up with a way to test a fifo, filo, lifo, etc. series of stacks/heaps/etc. with pointers in them that can automatically erase functions that are already sovled;

having some kind of click/touch/keystroke/etc. queue might make sense here as well.// perhaps thats what inpututil is for??
In short I would like the backend to have room for what would basically be a texturemap/imagemap that could be drawn from as well as a basic sound library that can perhaps add some basic effects to
each of the sounds?... as well as like do some stuff with regards to 'distance';


-----

### UI_Object


```
type UI_Object interface {
	// Init0() //initialize void;
	Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error //--
	Init_Parents(Parent UI_Object) error                                                                              //--
	Draw(screen *ebiten.Image) error                                                                                  //--
	Redraw()                                                                                                          //--
	Update() error                                                                                                    //--
	Update_Unactive() error                                                                                           //

	Update_Any() (any, error) //
	Update_Ret_State_Redraw_Status() (uint8, bool, error)
	Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error)

	GetState() uint8                                                    //
	ToString() string                                                   //
	IsInit() bool                                                       //
	GetID() string                                                      //
	GetType() string                                                    //
	IsCursorInBounds() bool                                             //
	IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool //
	GetPosition_Int() (int, int)                                        //
	GetNumber_Children() int                                            //
	GetChild(index int) UI_Object
	AddChild(child UI_Object) error //
	RemoveChild(index int) error
	GetParent() UI_Object //
	HasParent() bool      //
	// getType() string //might want to change this output to like an int or something using a golang equivelant of an enum;
}
```

I made this because I wanted to have a GUI framework with some sense of *sleek modularity*; ***but*** I'm having some troubles with it

notably I think I would be improved if it was a set of 3-4 separate interfaces that would be smaller and 'more narrow' in their applications while also allowing for some overlap;
hypothetically something like this:

```
type UI_Object_Lite interface{
    Init(backend *UI_Backend,style *UI_Style, ) error
    GetPosition(mode int) int,int
    Draw(screen *ebiten.Image) error
    Update() error
    IsCursorInBounds() bool
    
    ToString()
    GetStatus() int (??) <--- should this be something else? thre has to be a good way to get these things to function.
}

type UI_Object_Base struct{
    Position coords.CoordInts{}//do I even want it like that? wouldn't I prefer like having two floats32? or 2x float64?s
    Dimensions coords.CoordInts{}//<-----Not sure about this either; think about it there's going to be a 'base image' most likely that is being drawn to for this object; 
                                        (so I'm not drawing every subcomponent to the main image/screen buffer but a hundred little buffers)
                                        if the object is resizable though; or has like a second pane this changes things a little bit doesn't it?
    //Alternatively alternatively there's 'Rectangle' from the 'img' package;
    IsVisible bool
    IsActive bool
    Parent 
    //------------------------
}

```



```
type UI_Object_Style struct {
	LabelColor       color.Color
	PanelColor       color.Color
	BorderColor      color.Color
	ButtonColor0     []color.Color
	ButtonColor1     []color.Color
	BorderThickness  float32
	TextColor        []color.Color
	Internal_Margins [4]uint8
	TextSizes        []int
	TextAlignMode    int
}

```
apparently the syntax for a const as an enum replacement is something like

```
type Direction int
const(
    Up=Direction(iota)
    Down
    Right
    Left
)

```

### some other thoughts;

change 'Parent' (pointer) to "Context"; because these things will exist in a context;

fonts might need to be hadled through some kind of 'Rich Text' struct or interface that allows them to be manipulated to have 'italics/bolds/etc.';



## Interfaces that might make sense;

Requirements: UI-Objects;


--> tell parent that they need to refresh;
--> tell parent when they are minimized;
--> have a means by which to self-delete;
--> have a means by which to pass data back and forth; said data not being necessarily the same across all forms of ui-object;


how will a 'text' prompting window work? like will it require it to always exist but only sometimes be active? a simpler solution than having it be something that can just pop up but one that is also a bit... subpar??

will there need to be some kind of init?

UI_Init_Options?

x-mode buttons; with options to 'loop' or 'back and fourth', to have lists that cycle ?

## Backends;


```
type UI_Object_Type int
const(
	barebones = iota //or something like that
	button
	text field

)UI_Object_Type

type Basic_UI_Object struct{
	//blah blah blah...
	backend *UI_Backend
	layer uint8
	ui_id string
	type UI_Object_Type
}

type UI_Object interface{

}
```
The ulitmate 'backend' would include some kind of event system and event 'router'; 
```
Game_Backend{
	User_Interface []UI_Object //this is just a big list of pointers
	Audio_System //Some kind of audio system
	Texturemap tilemap //
	UI_Styles//--< this should include several default sizes for windows, buttons, 
	Typefaces//--< Ideally this should include means by which to get variations of a single font (so italics, BOLD text)
	UI_QUEUE func()
	touch_event_queue
	click_event_queue
	//--->
}

```
### click-touch handler;
 basically there needs to be somethng that can handle touch events in the UI where the UI will overlap (such as drop-down menus among other things)..
 this cannot be/should-not-be handled through 'load order' alone; though that's a solution that isn't ideal either;

 so I have a 'button_panel' with a dropdown menu;
 this dropdown menu extends outside the button window over another button_panel; and more importantly over another button;
 who has priority here? how is that decided?

 a priority queue might be a measure but how to implement that?
 --> idea; implement it through having a 'layer' integer; that is used to sort/resort the 'queue' and can be used to lock down the other objects being displayed;
 --> thus allowing for some kind of 'tag-out/lock out' system for lack of a better/catchy phrase.
 -->Problem how to implement this?
 I don't want a hundred bespoke little things like 'button group numbers' or whatever for every kind of object interface.
 I want
 {dropdownmenu.active:true}-->{dropdownmenu.parent.setpriority}-->{dropdownmenu.parent.parent.setpriority}

 perhaps having some kind of 'priority container' with a mandatory 'priority container cleared'.. this might also solve the problem of overlapping when not intented;

 //then again teh problems of overlapping would be solved by having 'layers' and having some kind of big draw queue;
// a system that was simultaniously like LIFO for drawing and FIFO for events/click-handling would be good here..
//-->

```
//type Game_Event interface? struct?
type Game_Event struct{
	timestamp_id//not sure what data for this..

}
```
having a queue of these might make some stuff slightly better.. but it seems like an awful lot of work infrastructurally for what is going to be fairly minor overall

