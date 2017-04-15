package rubber

type Rubber struct {
	Name		string	`db:"name"`
	Speed 		float32	`db:"speed"`
	Spin 		float32	`db:"spin"`
	Control 	float32	`db:"control"`
	Tackiness	float32	`db:"tackiness"`
	Weight		float32	`db:"weight"`
	SpongeHardness	float32	`db:"sponge_hardness"`
	Gears		float32	`db:"gears"`
	ThrowAngle	float32	`db:"throw_angle"`
	Consistency	float32	`db:"consistency"`
	Durability	float32	`db:"durability"`
}