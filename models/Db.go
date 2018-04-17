/**
 * @Desc
 * @author zjhfyq
 * @data 2018/3/29 15:09.
 */
package models

type DbConfig struct {
	DataSourceName string
	DriverName     string
	MaxIdleConns   int    
	MaxOpenConns   int
}
