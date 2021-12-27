package apps

import "github.com/luke513009828/crawlab-core/interfaces"

type App interface {
	Init()
	Start()
	Wait()
	Stop()
}

type NodeApp interface {
	App
	interfaces.WithConfigPath
	SetGrpcAddress(address interfaces.Address)
}

type MasterApp interface {
	NodeApp
	SetRunOnMaster(ok bool)
}

type WorkerApp interface {
	NodeApp
}
