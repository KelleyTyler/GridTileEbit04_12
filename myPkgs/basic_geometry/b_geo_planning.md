# **Basic Geometry: PLANNING**

## Intro:
This is a subpackage; possibly a future package on it's own that will have a set of things that I like and want for future use;

Ideally it will include multiple parts some of which will be modular some of which will not be;

This file is for marking down my thoughts or intentions for possible future developments.

---
## Coords:
- ### Coords
     Coordints needs to have some variations; ideally;
    - it requires uint (and uint8,uint16,etc.)
    - as well as just regular signed int8, int32 variations
    - it should also "play nice" with possible 'coordfloats'
    - in addition there should be 3-dimensional coordInts;
    - each of these will require their own "lists" perhaps lists themselves should follow a more modular format to handle the stacks involved;
    -
    #### ***Coords as Possible Interface***
    The point here would be to allow for coords/coordints to be more easily made into lists without having the troubling amount of re-writing every bit of code;
	    this might also be suited by having generics;
	
- ### CoordList
    - these might be better off handling a ***coord*** -as an *interface*;
    - Remove From CoordList (basically a list of values that coordlist is to remove at once)
    - GetCount from CoordList (basically a list of ) 
    - to CoordInt Array
    - from CoordInt Array 
    - 
---
## Matrix
- ### To Add
	The **Matrix** needs some additions;
	- Intmatrix variations with values of uint8--uint64 being present (should've honestly made this at first)
		- reason: uint matrices might be a better use of space and reduce ram/performance requirements 
		( even if only slightly)
		- 
	- 3d+ Matrices; with the functions that work for em;
	-  
	- FloatMatrix (float32 and Float64)
		- this is a bit difficult to think of right now
	- RuneMatrix?
	placeholder
- ### Etc.
- ### Matrix Pathfinding: ---
	 Pathfinding is going to be a major problem here I feel;
	The last iteration of this project I struggled to get pathfinding working to any acceptable degree;
	This time around I'm not sure it's going to be much better;
	The main problem is that I can 'find a path' but it's not necessarily the 'shortest path';
- ### Matrix Changes:
     Rather than have changes immediately processed to the board it might make more sense to have them 'buffered' by a combination of things;
     Basically either a struct or a combination of a CoordList (where the change is to take place), and a separate value or 'request' that is used to actually make the change.
     - this would need to work with the matrix in question to make sure the changes occured at planned
     - changes: 
        these would be most simply a 'set_to'- type request; "IE: set point X to Y"; but there might need to be more complicated sets of requests in the future depending on the nature of future development;
     
         
- ### ETC.;

---
## Some Other Thoughts

- should move things like 'integer_matrix_ebiten' to some kind of interface;

- 

---
 ## Pathfinding (cont.)
 Pathfinding should follow the ***Dijkstra’s Algorithm*** or the ***A\* Algorithm***.
 this means that we need a few things for a "node" in a linked list
   ### Node variables
 
 - **Starting Point**  CoordInts
 - **Target Point**	CoordInts
 - **Current position**	CoordInts
 - **Previous** a '\*Node'
 - **Next**(???) a '\*Node'
	 I question myself about this and whether or not I want to try and make this a ***[]\*Node{}*** 
	 (an array of node pointers).. the argument for it would be that it would allow for an easier selection 
 - **Distance-To-Start** float64
 - **Distance-To-End** float64
 - **Sum of Distances**	float64 
 - **Movement-Cost-To-Start** likely an integer
	  This would depend on a lot of factors most notably if I bother to implement having tiles that require additional movement costs to cross them; this would track the cost and be triggered by going up the list to the parent;
 - **Movement-Cost-To-End** likely an integer
	 Movement cost to end 


### Node functions 
In addition this *Node Object is going to need to have some functions; perhaps I should utilize the 'list' of collections/list for this as that might make it simpler; with prev/next offloaded to that; but disregarding 'linked list' attributes (where possible)
- **(node Node) Get Position bool** coordints
- **(node Node) OccupiesSamePosition(b \*Node)**   (bool)
- **(nodeA Node) Compare(nodeB \*Node)** ~~[4]bool~~ (samePosition, sameParent bool, costToStart, costToEnd uint8)
	- This would check if not only if the node in question occupied the same position but if it had the same pointers and same distance from start; it would return an array of bools
    - samePosition, sameParent (these are pretty self explanitory)
    - costToStart, costToEnd: these are a little more involved and will have a series of values that might just be better stored as const/enum-type arrangements;
        - (costToStart and costToEnd) = 0: equal
        - (costToStart and costToEnd) = 1: (nodeA) is more
        - (costToStart and costToEnd) = 2: (NodeB) is more
    - In theory we could have this reduced to like a 'byte'
        - [0,0,0,0,0,0,0,0]
        - this means we might be able to 'squeeze' two 3bit unsigned ints out of it; and still have room for two other bools!
        - this sounds hairbrained and likely to not be worth the time to implement however; 
Pathfinding Nodes would have to also have some kind of basic array structure that can handle some basic functions:
- Sorting:
    - sorting based on cost_to_start;
    - sorting based on cost_to
- Find if there is a duplicate entry, remove it
- Find if there is a near duplicate entry (IE: two nodes occupying the same position, this needs to be resolved)
    - resolving these shared things would need to depend on context however;
- set next/'child' pointers up and down the list; (this honestly might not be needed)
- Pop from 'front'
- push to back
- swap 2 nodes positions
- remove by index
- remove by position 
- remove by parent;

### Pathfinding Algorithm;
>  setup:
>   - **OpenList**  []*Node
>   - **ClosedList** []*Node
>   - **BlockedList** []*Node
>   - **Startpoint** coordint
>   - **Endpoint** coordint
>   - **IntMatrix** IntegerMatrix2d

--- sub functions:
  - node_Get_Neighbors_4_filtered(node \*Node, imat IntMatrix, filters []int, margin [4]int) (neighbors [4]\*Node, isvalid [4]bool)
   
    I hate how complicated these things quickly get; 
 - node_List_Sort_List

 - BlockedList manager
 - 

 #### actual algorithm 
    func Pathfind(start, target coords.CoordInts, imat mat.IntMatrix, ... , failconditions int) (path_is_found bool, path_out coords.CoordList){
        //setup
        path_is_found =false
        path_out = make(coords.CoordList,0)
        temp_Nodes: make([]*Node,0)
        //------
        //phase 1: ---> get an array of []nodes, or even a list of []nodes of potential paths (if they exist)
        path_is_found, temp_Nodes:= Pathfind_Phase1(start, target, imat, ...)
        //------
        //phase 2: ---> sort through paths and find the shortest one;
        if(path_is_found){
            path_out = Pathfind_Phase2A(tempNodes, start, target, imat, ... )
        }else{
            //reasoning it might be good to gets a set of paths 
            switch(failconditions){
                default:
                path_out = Pathfind_Phase2B(tempNodes, start,target,imat,...)
            }
        }

        return path_is_found, path_out 
    }


    func Pathfind_Phase1 (start, target coords.CoordInts, imat mat.IntMatrix, ... ) (path_is_found bool, PotentialPaths []*ImatNode){
        
        path_is_found= false
        OpenList:=make([]*ImatNode,0)
        ClosedList:=make([]*ImatNode,0)
        BlockedList:=make([]*ImatNode,0)
        var max_fails int = 100
        var curr_fails int =0
        var isFinished bool = false
        var err error = nil

        //beware pseudocode 
        //ClosedList.ClosedList = ClosedList.append(Get_New_Node(start, start, end))//*pseudocode or not I'm not sure I even want this.
        OpenList.append 
        for !isFinished{
            Openlist,ClosedList,BlockedList,isFinished,curr_fails,err = Pathfind_Phase1_Tick(Openlist, ClosedList, BlockedList,imat, isFinished,curr_fails, max_fails, imat,walls []int)
            if err!=nil{
                log.fatal(fmt.Errorf("Pathfinding Error"))
            }
        }   

        return potentialPaths
    }
    //-- this is the actual tick by tick process;
    func Pathfind_Phase1_Tick(start,end coords.CoordInts, openlist,closedlist,blockedlist []*ImatNode, pathfound bool,curr_fails, max_fails int, imat mat.IntegerMatrix2d, walls []int ) (oList,cList,bList []*ImatNode, pfound bool,fails int, err error){ //<--unsure if these are pass by reference or not 
        oList=make([]*ImatNode,len(openlist))
        cList=make([]*ImatNode,len(closedlist))
        bList=make([]*ImatNode,len(blockedlist))
        copy(oList, openlist)
        copy(cList, closedlist)
        copy(bList, blockedlist)
        pfound = pathfound
        fails = curr_fails
        pfound_phase2 = false (the path needs to be found more)
        //---------------------- Prep Done
        var pointQ *Node
        for len(oList)>0{
            pointQ, oList, err:= pseudocode.PopFromNodeArFront(oList)
            if err!=nil{
                return oList, cList, bList, false, err
            }
            if pointQ!=nil{
                if pointQ.Position.EqualTo(start){
                    //ending sequence
                    pFound=true
                }else if imat.IsValid(pointQ.position) && [value on the point is valid]{   
                    temp_successors := node_get_Neighbors_4(pointQ,imat)
                    //filter temp_successors into openlist 
                    //filter pointQ into closedlist comparing it with any overlap and resolving the contradiction;
                }


                oList= oList.sortByFValue()//f_value being sum of distance to and from
                oList= oList.removeDuplicates(nodesortopts_favorsmallestF)
                oList,cList,bList,err= Pathfind_Phase1_Tick_Blocked_List_Manager(oList,cList,bList)
                if err!=nil{
                    return oList, cList, bList, false, err
                }
            }


            
        }

        return oList,cList,bList, pfound, nil
    }
    //--this is to clean up and add stuff to the blockedlist
    func Pathfind_Phase1_Tick_Blocked_List_Manager(openlist,closedlist,blockedlist []*ImatNode) (oList,cList,bList []*ImatNode, error){ //<--unsure if these are pass by reference or not 
        oList=make([]*ImatNode,len(openlist))
        cList=make([]*ImatNode,len(closedlist))
        bList=make([]*ImatNode,len(blockedlist))
        copy(oList, openlist)
        copy(cList, closedlist)
        copy(bList, blockedlist)
        
        //---------------------- Prep Done
        for len(oList)>0{
        


        }

        return oList,cList,bList, nil
    }


## Some Other Pathfinding Thoughts