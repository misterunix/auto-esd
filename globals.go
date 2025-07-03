package main

type stable struct {
	Seed             int64    // Seed is the random seed used for image generation
	Width            int      // Width is the width of the image
	Height           int      // Height is the height of the image
	Prompt           string   // Prompt is the main prompt for the image generation
	Nprompt          string   // Nprompt is the negative prompt
	Model            string   // Model is the name of the model used for generation
	Sampler          string   // Sampler is the name of the sampler used for generation
	Steps            int      // Steps is the number of steps for the generation
	DateTime         string   // DateTime is the date and time of the generation and filename
	SmallImage       string   // SmallImage is the path and filename to the small image
	LargeImage       string   // LargeImage is the path and filename to the large image
	SmallImagePython string   // python script filename
	LargeImagePython string   // python script filename
	ModelsDir        string   // ModelsDir is the path to the models directory
	Models           []string // Models is the list of models available for generation
}
