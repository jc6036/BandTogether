package event_controller

import (
	"github.com/gin-gonic/gin"
)

func GetUserEvents(c *gin.Context) {
	c.JSON(200, GetEvents())
}

func GetEvents() []gin.H {
	return []gin.H{
		{
			"id":          "evt_rockmill",
			"title":       "Full-band rehearsal @ Rockmill Studio",
			"description": "Tighten the bridge on *Neon Skyline* and run the set twice. Bring in-ears.\n\n- Arrive 10 early for line check\n- Metronome at 106 bpm for **Skyline**\n- New harmony on chorus 2",
			"date":        "September 25, 2025",
			"start":       "19:00",
			"end":         "21:30",
			"users": []gin.H{
				{"id": "u1", "name": "Ava", "avatar": "https://i.pravatar.cc/64?img=32"},
				{"id": "u2", "name": "Miguel", "avatar": "https://i.pravatar.cc/64?img=5"},
				{"id": "u3", "name": "Sam", "avatar": "https://i.pravatar.cc/64?img=15"},
				{"id": "testid", "name": "Jesse Cabell", "avatar": "https://images.unsplash.com/photo-1502685104226-ee32379fefbe?w=256&q=80&auto=format&fit=crop"},
			},
			"createdAt": "September 1, 2025",
		},
		{
			"id":          "evt_cafe",
			"title":       "Acoustic duo – Blue Finch Cafe",
			"description": "2×45 sets. Bring DI, capo, merch square reader. House provides PA.",
			"date":        "September 22, 2025",
			"start":       "18:30",
			"end":         "20:15",
			"users": []gin.H{
				{"id": "u5", "name": "Riley", "avatar": "https://i.pravatar.cc/64?img=48"},
				{"id": "testid", "name": "Jesse Cabell", "avatar": "https://images.unsplash.com/photo-1502685104226-ee32379fefbe?w=256&q=80&auto=format&fit=crop"},
			},
			"createdAt": "September 1, 2025",
		},
		{
			"id":          "evt_fest",
			"title":       "Riverlights Festival - Main Stage",
			"description": "30-min changeover. City backline available. Parking passes in drive.",
			"date":        "September 26, 2025",
			"start":       "16:00",
			"end":         "17:00",
			"users": []gin.H{
				{"id": "u6", "name": "Taylor", "avatar": "https://i.pravatar.cc/64?img=21"},
				{"id": "u7", "name": "Jordan", "avatar": "https://i.pravatar.cc/64?img=11"},
				{"id": "testid", "name": "Jesse Cabell", "avatar": "https://images.unsplash.com/photo-1502685104226-ee32379fefbe?w=256&q=80&auto=format&fit=crop"},
			},
			"createdAt": "September 1, 2025",
		},
	}
}
