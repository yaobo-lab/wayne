package model

type configModel struct{}

type ConfigKey string

const (
	ConfigKeyTitile              ConfigKey = "system.title"
	ConfigKeyImageNamePrefix     ConfigKey = "system.image-prefix"
	ConfigKeyAffinity            ConfigKey = "system.affinity"
	ConfigKeyMonitorUri          ConfigKey = "system.monitor-uri"
	ConfigKeyApiNameGenerateRule ConfigKey = "system.api-name-generate-rule"
	ConfigKeyOauth2Title         ConfigKey = "system.oauth2-title"
	TableNameConfig                        = "config"
)

type Config struct {
	Id    int64     `orm:"auto" json:"id,omitempty"`
	Name  ConfigKey `orm:"size(256)" json:"name,omitempty"`
	Value string    `orm:"type(text)" json:"value,omitempty"`
}

func (*Config) TableName() string {
	return TableNameConfig
}

func (*configModel) Add(m *Config) (id int64, err error) {
	id, err = Ormer().Insert(m)
	if err != nil {
		return
	}
	return
}

func (*configModel) GetById(id int64) (v *Config, err error) {
	v = &Config{Id: id}

	if err = Ormer().Read(v); err != nil {
		return nil, err
	}
	return v, err
}

func (*configModel) UpdateById(m *Config) (err error) {
	v := Config{Id: m.Id}
	if err = Ormer().Read(&v); err == nil {
		_, err = Ormer().Update(m)
		return err
	}
	return
}

func (*configModel) DeleteById(id int64) (err error) {
	v := Config{Id: id}
	if err = Ormer().Read(&v); err == nil {
		_, err = Ormer().Delete(&v)
		return err
	}
	return
}
