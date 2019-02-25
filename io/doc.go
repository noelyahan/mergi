/*
Package io implements Importer, Exporter interfaces for library inputs and outputs

Loader pkg supports

	FileExporter
	FileImporter
	URLImporter
	AnimationExporter

Loader pkg can be extends to more importers and exporters
	ex: KafkaImporter, KafkaExporter, MQTT, WebSockets, Serial etc..
Create your own io and let us know.
*/
package io // import "github.com/noelyahan/mergi/io"
