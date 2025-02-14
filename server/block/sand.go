 

 package block 

  

 import ( 

         "github.com/Pocketminer92/magic-alpaca/server/block/cube" 

         "github.com/Pocketminer92/magic-alpaca/server/world" 

 ) 

  

 // Sand is a block affected by gravity. It can come in a red variant. 

 type Sand struct { 

         gravityAffected 

         solid 

         snare 

  

         // Red toggles the red sand variant. 

         Red bool 

 } 

  

 // SoilFor ... 

 func (s Sand) SoilFor(block world.Block) bool { 

         _, ok := block.(DeadBush) 

         return ok 

 } 

  

 // NeighbourUpdateTick ... 

 func (s Sand) NeighbourUpdateTick(pos, _ cube.Pos, w *world.World) { 

         s.fall(s, pos, w) 

 } 

  

 // BreakInfo ... 

 func (s Sand) BreakInfo() BreakInfo { 

         return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(s)) 

 } 

  

 // EncodeItem ... 

 func (s Sand) EncodeItem() (name string, meta int16) { 

         if s.Red { 

                 return "minecraft:sand", 1 

         } 

         return "minecraft:sand", 0 

 } 

  

 // EncodeBlock ... 

 func (s Sand) EncodeBlock() (string, map[string]any) { 

         if s.Red { 

                 return "minecraft:sand", map[string]any{"sand_type": "red"} 

         } 

         return "minecraft:sand", map[string]any{"sand_type": "normal"} 

 }
