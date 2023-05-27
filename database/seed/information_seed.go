package seed

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func CreateInformation() []*ei.Information {
	information := []*ei.Information{
		{
			Title:           "Title Information 1",
			PhotoContentUrl: "Photo Content URL 1",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			StatusId:        1,
		},
		{
			Title:           "Title Information 2",
			PhotoContentUrl: "Photo Content URL 3",
			Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.",
			StatusId:        2,
		},
	}
	return information
}
