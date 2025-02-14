 

 package block 

  

 import ( 

         "github.com/Pocketminer92/magic-alpaca/server/item" 

 ) 

  

 // RawIronBlock is a raw metal block equivalent to nine raw iron. 

 type RawIronBlock struct { 

         solid 

         bassDrum 

 } 

  

 // BreakInfo ... 

 func (r RawIronBlock) BreakInfo() BreakInfo { 

         return newBreakInfo(5, func(t item.Tool) bool { 

                 return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel 

         }, pickaxeEffective, oneOf(r)) 

 } 

  

 // EncodeItem ... 

 func (RawIronBlock) EncodeItem() (name string, meta int16) { 

         return "minecraft:raw_iron_block", 0 

 } 

  

 // EncodeBlock ... 

 func (RawIronBlock) EncodeBlock() (string, map[string]any) { 

         return "minecraft:raw_iron_block", nil 

 }
