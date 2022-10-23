package config

type DataMock struct {
	Interval      int
	MockState     bool
	WriteRedisMax int
}

var DataMockConfig = new(DataMock)
