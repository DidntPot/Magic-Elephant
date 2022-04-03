 

 package model 

  

 import ( 

         "github.com/Pocketminer92/magic-alpaca/server/block/cube" 

         "github.com/Pocketminer92/magic-alpaca/server/entity/physics" 

         "github.com/Pocketminer92/magic-alpaca/server/world" 

         "github.com/go-gl/mathgl/mgl64" 

 ) 

  

 // Carpet is a model for carpet-like extremely thin blocks. 

 type Carpet struct{} 

  

 // AABB returns a flat AABB with a width of 0.0625. 

 func (Carpet) AABB(cube.Pos, *world.World) []physics.AABB { 

         return []physics.AABB{physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{1, 0.0625, 1})} 

 } 

  

 // FaceSolid always returns false. 

 func (Carpet) FaceSolid(cube.Pos, cube.Face, *world.World) bool { 

         return false 

 }
