package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func CreateInformation() []*ei.Information {
	information := []*ei.Information{
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 1",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 2",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 3",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 4",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 5",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 6",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 7",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 8",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 9",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 10",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 11",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 12",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 13",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Terbit",
		},
		{
			InformationId:   randomid.GenerateRandomNumber(),
			Title:           "Title Information 14",
			PhotoContentUrl: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			Status:          "Draft",
		},
	}
	return information
}
