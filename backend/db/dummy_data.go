package db

import (
	"real-time-forum/pkg/model"
	"time"
)

var Users = []model.User{
	{1, "Kaspar", "Pärtel", 22, "male", "KasparP", "kasparjohannes@gmail.com", "123456", time.Now()},
	{2, "Erkki", "Tikk", 37, "male", "ErkkiT", "erkki.tikk@gmail.com", "123456", time.Now()},
}

var DummyPosts = []model.Post{
	{1,
		"Random postitus 1",
		"Vivamus libero tortor, mattis non nisl a, malesuada dignissim erat. Suspendisse mollis euismod dui in tristique. Fusce quis efficitur lectus. Aliquam lacinia, turpis ac pharetra bibendum, metus risus sodales eros, at rhoncus nunc felis at ex. Praesent fermentum ligula vel nisi vehicula aliquam.",
		Users[0],
		"image.jpg",
		5,
		6,
		2,
		time.Now(),
	},
	{2,
		"Random postitus 2",
		"Vivamus libero tortor, mattis non nisl a, malesuada dignissim erat. Suspendisse mollis euismod dui in tristique. Fusce quis efficitur lectus. Aliquam lacinia, turpis ac pharetra bibendum, metus risus sodales eros, at rhoncus nunc felis at ex. Praesent fermentum ligula vel nisi vehicula aliquam.",
		Users[1],
		"image.jpg",
		3,
		0,
		1,
		time.Now(),
	},
	{3,
		"Random postitus 3",
		"Vivamus libero tortor, mattis non nisl a, malesuada dignissim erat. Suspendisse mollis euismod dui in tristique. Fusce quis efficitur lectus. Aliquam lacinia, turpis ac pharetra bibendum, metus risus sodales eros, at rhoncus nunc felis at ex. Praesent fermentum ligula vel nisi vehicula aliquam.",
		Users[0],
		"image.jpg",
		0,
		0,
		0,
		time.Now(),
	},
	{4,
		"Random postitus 4",
		"Vivamus libero tortor, mattis non nisl a, malesuada dignissim erat. Suspendisse mollis euismod dui in tristique. Fusce quis efficitur lectus. Aliquam lacinia, turpis ac pharetra bibendum, metus risus sodales eros, at rhoncus nunc felis at ex. Praesent fermentum ligula vel nisi vehicula aliquam.",
		Users[1],
		"image.jpg",
		12,
		3,
		7,
		time.Now(),
	},
	{5,
		"Random postitus 5",
		"Vivamus libero tortor, mattis non nisl a, malesuada dignissim erat. Suspendisse mollis euismod dui in tristique. Fusce quis efficitur lectus. Aliquam lacinia, turpis ac pharetra bibendum, metus risus sodales eros, at rhoncus nunc felis at ex. Praesent fermentum ligula vel nisi vehicula aliquam.",
		Users[0],
		"image.jpg",
		3,
		10,
		8,
		time.Now(),
	},
}
