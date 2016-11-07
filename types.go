package nextbus

type Agency struct {
	Title       string
	RegionTitle string
	Tag         string
}

type Route struct {
	Title      string
	ShortTitle string
	Tag        string
}

type RouteConfig struct {
	Title         string
	ShortTitle    string
	Tag           string
	LatMin        string
	LonMin        string
	LatMax        string
	LonMax        string
	Color         string
	OppositeColor string
	Stop          []Stop
	Direction     []struct {
		Stop     []Stop
		Title    string
		UseForUI string
		Tag      string
		Name     string
	}
	Path []struct {
		Point []struct {
			Lat string
			Lon string
		}
	}
}

type Stop struct {
	Tag        string
	Title      string
	ShortTitle string
	Lat        string
	Lon        string
	StopId     string
}

type Predictions struct {
	AgencyTitle                  string
	RouteTag                     string
	RouteTitle                   string
	StopTag                      string
	StopTitle                    string
	DirTitleBecauseNoPredictions string
	Direction                    struct {
		Title      string
		Prediction []Prediction
	}
	Message []struct {
		Text     string
		Priority string
	}
}

type Prediction struct {
	Seconds           string
	Minutes           string
	EpochTime         string
	IsDeparture       string
	DirTag            string
	Block             string
	TripTag           string
	Branch            string
	AffectedByLayover string
	IsScheduleBased   string
	Delayed           string
	Vehicle           string
}

type Schedule struct {
	ServiceClass string
	Title        string
	Tr           []Tr
	Direction    string
	Tag          string
	Header       struct {
		Stop     []struct {
			Content string
			Tag     string
		}
	}
	ScheduleClass string
}

type Tr struct {
	Stop []struct {
		Content   string
		Tag       string
		EpochTime string
	}
	BlockId string
}