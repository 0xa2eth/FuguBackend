package password

import (
	"fmt"
	"math/rand"
	"time"
)

// 鱼类种类
var fishTypes = []string{
	"Goldfish", "Clownfish", "Salmon", "Tuna", "Trout", "Cod", "Snapper", "Bass", "Sardine", "Mackerel",
	"Swordfish", "Catfish", "Guppy", "Angelfish", "Piranha", "Seahorse", "Marlin", "Mahi-mahi", "Flounder",
	"Carp", "Haddock", "Herring", "Perch", "Tilapia", "Pike", "Anchovy", "Halibut", "Sole", "Barracuda",
	"Eel", "Jellyfish", "Starfish", "Stingray", "Triggerfish", "Wrasse", "Pufferfish", "Lionfish", "Archerfish",
	"Butterflyfish", "Parrotfish", "Damselfish", "Surgeonfish", "Dory", "Blobfish", "Bluefish", "Grouper",
	"Drumfish", "Wolffish", "Monkfish", "Redfish", "Turbot", "Snakehead", "Mudskipper", "Mudfish", "Garfish",
	"Fluke", "Sturgeon", "Shrimpfish", "Scampi", "Cuttlefish", "Squid", "Octopus", "Krill", "Lobster",
	"Crab", "Crawfish", "Shrimp", "Prawn", "Clam", "Mussel", "Oyster", "Scallop", "Abalone", "Barnacle",
	"Conch", "Whelk", "Snail", "Jellyfish", "Sea cucumber", "Sea urchin", "Sand dollar", "Sea star",
	"Coral", "Sponge", "Anemone", "Sea fan", "Sea horse", "Sea turtle", "Dolphin", "Shark", "Whale",
	"Marlin", "Swordfish", "Tuna", "Ray", "Barracuda", "Moray eel", "Pufferfish", "Anglerfish",
}

// 形容词
var adjectives = []string{
	"Attractive", "Appealing", "Alluring", "Beautiful", "Captivating", "Charming", "Elegant", "Exquisite", "Fascinating", "Gorgeous",
	"Handsome", "Irresistible", "Lovely", "Mesmerizing", "Radiant", "Ravishing", "Seductive", "Stunning", "Breathtaking", "Enchanting",
	"Graceful", "Delightful", "Pretty", "Striking", "Sophisticated", "Magnetic", "Adorable", "Charming", "Cute", "Dashing",
	"Elegant", "Fashionable", "Glamorous", "Hunky", "Impeccable", "Majestic", "Polished", "Refined", "Sleek", "Stylish",
	"Suave", "Trendy", "Upscale", "Voluptuous", "Yummy", "Alluring", "Bewitching", "Sensual", "Provocative", "Passionate",
	"Hot", "Sexy", "Seductive", "Sizzling", "Sultry", "Steamy", "Desirable", "Luscious", "Voluptuous", "Tempting",
	"Appealing", "Charismatic", "Charming", "Enchanting", "Enticing", "Fascinating", "Irresistible", "Magnetic", "Alluring", "Seductive",
	"Spellbinding", "Stunning", "Sensual", "Sizzling", "Radiant", "Captivating", "Elegant", "Glamorous", "Exquisite", "Attractive",
	"Charming", "Handsome", "Gorgeous", "Breathtaking", "Ravishing", "Lovely", "Enchanting", "Mesmerizing", "Striking", "Fascinating",
	"Stunning", "Dazzling", "Alluring", "Seductive", "Hot", "Sexy", "Sultry", "Voluptuous", "Passionate", "Desirable",
}

// GetRandomFishName 返回随机的鱼类描述，包括种类和形容词的组合
func GetRandomFishName() string {
	rand.Seed(time.Now().UnixNano())
	fishType := fishTypes[rand.Intn(len(fishTypes))]
	adj := adjectives[rand.Intn(len(adjectives))]
	str := GenRandomStr(string(rand.Int63n(99999)), 3)
	return fmt.Sprintf("%s%s-", adj, fishType) + str
}
