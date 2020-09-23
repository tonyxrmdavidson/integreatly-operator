package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Link = map[string]string{
	"":     "Represents a standard link that could be generated in HTML",
	"text": "text is the display text for the link",
	"href": "href is the absolute secure URL for the link (must use https)",
}

func (Link) SwaggerDoc() map[string]string {
	return map_Link
}

var map_CLIDownloadLink = map[string]string{
	"text": "text is the display text for the link",
	"href": "href is the absolute secure URL for the link (must use https)",
}

func (CLIDownloadLink) SwaggerDoc() map[string]string {
	return map_CLIDownloadLink
}

var map_ConsoleCLIDownload = map[string]string{
	"": "ConsoleCLIDownload is an extension for configuring openshift web console command line interface (CLI) downloads.",
}

func (ConsoleCLIDownload) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownload
}

var map_ConsoleCLIDownloadSpec = map[string]string{
	"":            "ConsoleCLIDownloadSpec is the desired cli download configuration.",
	"displayName": "displayName is the display name of the CLI download.",
	"description": "description is the description of the CLI download (can include markdown).",
	"links":       "links is a list of objects that provide CLI download link details.",
}

func (ConsoleCLIDownloadSpec) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownloadSpec
}

var map_ConsoleExternalLogLink = map[string]string{
	"": "ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.",
}

func (ConsoleExternalLogLink) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLink
}

var map_ConsoleExternalLogLinkSpec = map[string]string{
	"":                "ConsoleExternalLogLinkSpec is the desired log link configuration. The log link will appear on the logs tab of the pod details page.",
	"text":            "text is the display text for the link",
	"hrefTemplate":    "hrefTemplate is an absolute secure URL (must use https) for the log link including variables to be replaced. Variables are specified in the URL with the format ${variableName}, for instance, ${containerName} and will be replaced with the corresponding values from the resource. Resource is a pod. Supported variables are: - ${resourceName} - name of the resource which containes the logs - ${resourceUID} - UID of the resource which contains the logs\n              - e.g. `11111111-2222-3333-4444-555555555555`\n- ${containerName} - name of the resource's container that contains the logs - ${resourceNamespace} - namespace of the resource that contains the logs - ${resourceNamespaceUID} - namespace UID of the resource that contains the logs - ${podLabels} - JSON representation of labels matching the pod with the logs\n            - e.g. `{\"key1\":\"value1\",\"key2\":\"value2\"}`\n\ne.g., https://example.com/logs?resourceName=${resourceName}&containerName=${containerName}&resourceNamespace=${resourceNamespace}&podLabels=${podLabels}",
	"namespaceFilter": "namespaceFilter is a regular expression used to restrict a log link to a matching set of namespaces (e.g., `^openshift-`). The string is converted into a regular expression using the JavaScript RegExp constructor. If not specified, links will be displayed for all the namespaces.",
}

func (ConsoleExternalLogLinkSpec) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLinkSpec
}

var map_ApplicationMenuSpec = map[string]string{
	"":         "ApplicationMenuSpec is the specification of the desired section and icon used for the link in the application menu.",
	"section":  "section is the section of the application menu in which the link should appear. This can be any text that will appear as a subheading in the application menu dropdown. A new section will be created if the text does not match text of an existing section.",
	"imageURL": "imageUrl is the URL for the icon used in front of the link in the application menu. The URL must be an HTTPS URL or a Data URI. The image should be square and will be shown at 24x24 pixels.",
}

func (ApplicationMenuSpec) SwaggerDoc() map[string]string {
	return map_ApplicationMenuSpec
}

var map_ConsoleLink = map[string]string{
	"": "ConsoleLink is an extension for customizing OpenShift web console links.",
}

func (ConsoleLink) SwaggerDoc() map[string]string {
	return map_ConsoleLink
}

var map_ConsoleLinkSpec = map[string]string{
	"":                   "ConsoleLinkSpec is the desired console link configuration.",
	"location":           "location determines which location in the console the link will be appended to (ApplicationMenu, HelpMenu, UserMenu, NamespaceDashboard).",
	"applicationMenu":    "applicationMenu holds information about section and icon used for the link in the application menu, and it is applicable only when location is set to ApplicationMenu.",
	"namespaceDashboard": "namespaceDashboard holds information about namespaces in which the dashboard link should appear, and it is applicable only when location is set to NamespaceDashboard. If not specified, the link will appear in all namespaces.",
}

func (ConsoleLinkSpec) SwaggerDoc() map[string]string {
	return map_ConsoleLinkSpec
}

var map_NamespaceDashboardSpec = map[string]string{
	"":                  "NamespaceDashboardSpec is a specification of namespaces in which the dashboard link should appear. If both namespaces and namespaceSelector are specified, the link will appear in namespaces that match either",
	"namespaces":        "namespaces is an array of namespace names in which the dashboard link should appear.",
	"namespaceSelector": "namespaceSelector is used to select the Namespaces that should contain dashboard link by label. If the namespace labels match, dashboard link will be shown for the namespaces.",
}

func (NamespaceDashboardSpec) SwaggerDoc() map[string]string {
	return map_NamespaceDashboardSpec
}

var map_ConsoleNotification = map[string]string{
	"": "ConsoleNotification is the extension for configuring openshift web console notifications.",
}

func (ConsoleNotification) SwaggerDoc() map[string]string {
	return map_ConsoleNotification
}

var map_ConsoleNotificationSpec = map[string]string{
	"":                "ConsoleNotificationSpec is the desired console notification configuration.",
	"text":            "text is the visible text of the notification.",
	"location":        "location is the location of the notification in the console.",
	"link":            "link is an object that holds notification link details.",
	"color":           "color is the color of the text for the notification as CSS data type color.",
	"backgroundColor": "backgroundColor is the color of the background for the notification as CSS data type color.",
}

func (ConsoleNotificationSpec) SwaggerDoc() map[string]string {
	return map_ConsoleNotificationSpec
}

var map_ConsoleYAMLSample = map[string]string{
	"": "ConsoleYAMLSample is an extension for customizing OpenShift web console YAML samples.",
}

func (ConsoleYAMLSample) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSample
}

var map_ConsoleYAMLSampleSpec = map[string]string{
	"":               "ConsoleYAMLSampleSpec is the desired YAML sample configuration. Samples will appear with their descriptions in a samples sidebar when creating a resources in the web console.",
	"TargetResource": "targetResource contains apiVersion and kind of the resource YAML sample is representating.",
	"title":          "title of the YAML sample.",
	"description":    "description of the YAML sample.",
	"yaml":           "yaml is the YAML sample to display.",
}

func (ConsoleYAMLSampleSpec) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSampleSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
