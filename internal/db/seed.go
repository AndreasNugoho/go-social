package db

import (
	"context"
	"fmt"
	"log"

	"math/rand"

	"github.com/AndreasNugoho/go-social/internal/store"
)

var usernames = []string{
	"andre",
	"bob",
	"nugroho",
	"david",
	"sarah",
	"alex",
	"luna",
	"mike",
	"jane",
	"ryan",
	"emma",
	"john",
	"olivia",
	"chris",
	"mia",
	"alex",
	"ava",
	"daniel",
	"sophia",
	"matthew",
	"isabella",
	"william",
	"charlotte",
	"james",
	"amelia",
	"benjamin",
	"harper",
	"evelyn",
	"lucas",
	"abigail",
	"julia",
	"grace",
	"sebastian",
	"eleanor",
	"thomas",
	"ella",
	"lucas",
	"clara",
	"max",
	"lily",
	"harry",
	"hannah",
	"mason",
	"natalie",
	"noah",
	"madison",
	"jack",
	"scarlett",
	"leo",
}

var titles = []string{
	"Software Engineer",
	"Product Manager",
	"Data Scientist",
	"UX Designer",
	"DevOps Engineer",
	"Frontend Developer",
	"Backend Developer",
	"Machine Learning Engineer",
	"Database Administrator",
	"Security Analyst",
	"Systems Architect",
	"Technical Lead",
	"Cloud Engineer",
	"Quality Assurance Engineer",
	"Mobile Developer",
	"Full Stack Developer",
	"Business Analyst",
	"IT Support Specialist",
	"Network Administrator",
	"Solution Architect",
}

var contents = []string{
	"Welcome to our new platform! We're excited to launch this exciting update.",
	"This feature was built with performance and usability in mind.",
	"Our team has been working tirelessly to deliver this milestone.",
	"Check out the latest updates and what's new in version 2.0.",
	"User feedback has shaped this release, and we've made improvements accordingly.",
	"We're committed to continuous innovation and better user experiences.",
	"This update includes security enhancements and bug fixes.",
	"The new interface is cleaner, faster, and more intuitive.",
	"We've optimized backend processes for improved response times.",
	"Enjoy a smoother experience across all devices and platforms.",
	"Our documentation has been updated to help you get started quickly.",
	"We're excited to see how you use these new features in your workflow.",
	"This release marks a significant step forward in our roadmap.",
	"Stay tuned for more updates rolling out in the coming months.",
	"Your input matters — tell us what you think about the new design.",
	"Performance benchmarks show a 40% improvement in load times.",
	"We've added dark mode support for reduced eye strain.",
	"The mobile app now integrates seamlessly with desktop features.",
	"Multi-language support is now available for global users.",
	"Thank you for being part of our growing community of users.",
}

var tags = []string{
	"golang",
	"web-development",
	"api",
	"microservices",
	"cloud",
	"devops",
	"database",
	"security",
	"performance",
	"testing",
	"design-patterns",
	"frontend",
	"backend",
	"architecture",
	"containers",
	"docker",
	"kubernetes",
	"authentication",
	"authorization",
	"ci-cd",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "password",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				titles[rand.Intn(len(tags))],
				titles[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: contents[rand.Intn(len(contents))],
		}
	}

	return cms
}
