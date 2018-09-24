/*
Package loader implements Importer, Exporter interfaces for library inputs and outputs

Loader pkg supports

	FileExporter
	FileImporter
	URLImporter
	AnimationExporter

Loader pkg can be extends to more importers and exporters
	ex: KafkaImporter, KafkaExporter, MQTT, WebSockets, Serial etc..
Create your own loader and let us know.
*/
package loader // import "github.com/noelyahan/mergi/loader"
