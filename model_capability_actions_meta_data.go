/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9235
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// CapabilityActionsMetaData Metadata or constraints of various server actions supported in Intersight for 3rd Party servers. It is validated against the target provided  and the actions are allowed only upon successful validation.
type CapabilityActionsMetaData struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum firmware version of the server.
	MaxFwVersion *string `json:"MaxFwVersion,omitempty"`
	// Minimum firmware version of the server.
	MinFwVersion *string `json:"MinFwVersion,omitempty"`
	// Model of the server that supports power or storage operations.
	Model            *string  `json:"Model,omitempty"`
	SupportedActions []string `json:"SupportedActions,omitempty"`
	// The target type to which the metadata applies. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `CAPIC` - A Cloud Application Policy Infrastructure Controller instance. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `ACISwitch` - A platform type to support ACI Switches. * `NexusSwitch` - A platform type to support Cisco Nexus Switches. * `MDSSwitch` - A platform type to support Cisco MDS Switches. * `MDSDevice` - A platform type to support MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `RedfishServer` - A generic target type for servers that support Redfish. Current support is limited to managing HPE and Dell servers on Intersight. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `MySqlServer` - An instance of either Oracle MySQL Database or the open source MariaDB. * `OracleDatabaseServer` - The Oracle Server is a relational database management system that provides an open, comprehensive, and integrated approach to information management. * `IBMWebSphereApplicationServer` - WebSphere Application Server (WAS) is a software product that performs the role of a web application server. More specifically it is a software framework and middleware that hosts Java based web applications. * `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * `ApacheTomcatServer` - Apache Tomcat is a web container. It allows the users to run Servlet and JAVA Server Pages that are based on the web-applications. * `JavaVirtualMachine` - The Java Virtual Machine (JVM) is the runtime engine of the Java Platform, which allows any program written in Java or other language compiled into Java bytecode to run on any computer that has a native JVM. * `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - An Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - An Amazon web service billing target to retrieve billing information stored in S3 bucket. * `GoogleCloudPlatform` - Google Cloud Platform (GCP), offered by Google, is a suite of cloud computing services that runs on the same infrastructure that Google uses internally for its end-user products, such as Google Search, Gmail, Google Drive, and YouTube. Alongside a set of management tools, it provides a series of modular cloud services including computing, data storage, data analytics and machine learning. Google Cloud Platform provides infrastructure as a service, platform as a service, and serverless computing environments. * `GoogleCloudPlatformBilling` - Google Cloud Platform (GCP) offers flexible ways to set up and manage billing for your resources. A billing account is how a user pays for the resources being consumed. A billing account is associated with a method of payment and access is established using Cloud IAM roles. For a resource to be deployed in a project, the project has to be associated with a billing account. More than one project can be associated with a billing account. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
	TargetType *string `json:"TargetType,omitempty"`
	// Manufacturer of the server.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityActionsMetaData CapabilityActionsMetaData

// NewCapabilityActionsMetaData instantiates a new CapabilityActionsMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityActionsMetaData(classId string, objectType string) *CapabilityActionsMetaData {
	this := CapabilityActionsMetaData{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewCapabilityActionsMetaDataWithDefaults instantiates a new CapabilityActionsMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityActionsMetaDataWithDefaults() *CapabilityActionsMetaData {
	this := CapabilityActionsMetaData{}
	var classId string = "capability.ActionsMetaData"
	this.ClassId = classId
	var objectType string = "capability.ActionsMetaData"
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityActionsMetaData) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityActionsMetaData) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityActionsMetaData) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityActionsMetaData) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaxFwVersion returns the MaxFwVersion field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetMaxFwVersion() string {
	if o == nil || o.MaxFwVersion == nil {
		var ret string
		return ret
	}
	return *o.MaxFwVersion
}

// GetMaxFwVersionOk returns a tuple with the MaxFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetMaxFwVersionOk() (*string, bool) {
	if o == nil || o.MaxFwVersion == nil {
		return nil, false
	}
	return o.MaxFwVersion, true
}

// HasMaxFwVersion returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasMaxFwVersion() bool {
	if o != nil && o.MaxFwVersion != nil {
		return true
	}

	return false
}

// SetMaxFwVersion gets a reference to the given string and assigns it to the MaxFwVersion field.
func (o *CapabilityActionsMetaData) SetMaxFwVersion(v string) {
	o.MaxFwVersion = &v
}

// GetMinFwVersion returns the MinFwVersion field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetMinFwVersion() string {
	if o == nil || o.MinFwVersion == nil {
		var ret string
		return ret
	}
	return *o.MinFwVersion
}

// GetMinFwVersionOk returns a tuple with the MinFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetMinFwVersionOk() (*string, bool) {
	if o == nil || o.MinFwVersion == nil {
		return nil, false
	}
	return o.MinFwVersion, true
}

// HasMinFwVersion returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasMinFwVersion() bool {
	if o != nil && o.MinFwVersion != nil {
		return true
	}

	return false
}

// SetMinFwVersion gets a reference to the given string and assigns it to the MinFwVersion field.
func (o *CapabilityActionsMetaData) SetMinFwVersion(v string) {
	o.MinFwVersion = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityActionsMetaData) SetModel(v string) {
	o.Model = &v
}

// GetSupportedActions returns the SupportedActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityActionsMetaData) GetSupportedActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedActions
}

// GetSupportedActionsOk returns a tuple with the SupportedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityActionsMetaData) GetSupportedActionsOk() ([]string, bool) {
	if o == nil || o.SupportedActions == nil {
		return nil, false
	}
	return o.SupportedActions, true
}

// HasSupportedActions returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasSupportedActions() bool {
	if o != nil && o.SupportedActions != nil {
		return true
	}

	return false
}

// SetSupportedActions gets a reference to the given []string and assigns it to the SupportedActions field.
func (o *CapabilityActionsMetaData) SetSupportedActions(v []string) {
	o.SupportedActions = v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *CapabilityActionsMetaData) SetTargetType(v string) {
	o.TargetType = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityActionsMetaData) SetVendor(v string) {
	o.Vendor = &v
}

func (o CapabilityActionsMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaxFwVersion != nil {
		toSerialize["MaxFwVersion"] = o.MaxFwVersion
	}
	if o.MinFwVersion != nil {
		toSerialize["MinFwVersion"] = o.MinFwVersion
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.SupportedActions != nil {
		toSerialize["SupportedActions"] = o.SupportedActions
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityActionsMetaData) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityActionsMetaDataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum firmware version of the server.
		MaxFwVersion *string `json:"MaxFwVersion,omitempty"`
		// Minimum firmware version of the server.
		MinFwVersion *string `json:"MinFwVersion,omitempty"`
		// Model of the server that supports power or storage operations.
		Model            *string  `json:"Model,omitempty"`
		SupportedActions []string `json:"SupportedActions,omitempty"`
		// The target type to which the metadata applies. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `CAPIC` - A Cloud Application Policy Infrastructure Controller instance. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `ACISwitch` - A platform type to support ACI Switches. * `NexusSwitch` - A platform type to support Cisco Nexus Switches. * `MDSSwitch` - A platform type to support Cisco MDS Switches. * `MDSDevice` - A platform type to support MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `RedfishServer` - A generic target type for servers that support Redfish. Current support is limited to managing HPE and Dell servers on Intersight. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `MySqlServer` - An instance of either Oracle MySQL Database or the open source MariaDB. * `OracleDatabaseServer` - The Oracle Server is a relational database management system that provides an open, comprehensive, and integrated approach to information management. * `IBMWebSphereApplicationServer` - WebSphere Application Server (WAS) is a software product that performs the role of a web application server. More specifically it is a software framework and middleware that hosts Java based web applications. * `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * `ApacheTomcatServer` - Apache Tomcat is a web container. It allows the users to run Servlet and JAVA Server Pages that are based on the web-applications. * `JavaVirtualMachine` - The Java Virtual Machine (JVM) is the runtime engine of the Java Platform, which allows any program written in Java or other language compiled into Java bytecode to run on any computer that has a native JVM. * `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - An Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - An Amazon web service billing target to retrieve billing information stored in S3 bucket. * `GoogleCloudPlatform` - Google Cloud Platform (GCP), offered by Google, is a suite of cloud computing services that runs on the same infrastructure that Google uses internally for its end-user products, such as Google Search, Gmail, Google Drive, and YouTube. Alongside a set of management tools, it provides a series of modular cloud services including computing, data storage, data analytics and machine learning. Google Cloud Platform provides infrastructure as a service, platform as a service, and serverless computing environments. * `GoogleCloudPlatformBilling` - Google Cloud Platform (GCP) offers flexible ways to set up and manage billing for your resources. A billing account is how a user pays for the resources being consumed. A billing account is associated with a method of payment and access is established using Cloud IAM roles. For a resource to be deployed in a project, the project has to be associated with a billing account. More than one project can be associated with a billing account. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
		TargetType *string `json:"TargetType,omitempty"`
		// Manufacturer of the server.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varCapabilityActionsMetaDataWithoutEmbeddedStruct := CapabilityActionsMetaDataWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityActionsMetaDataWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityActionsMetaData := _CapabilityActionsMetaData{}
		varCapabilityActionsMetaData.ClassId = varCapabilityActionsMetaDataWithoutEmbeddedStruct.ClassId
		varCapabilityActionsMetaData.ObjectType = varCapabilityActionsMetaDataWithoutEmbeddedStruct.ObjectType
		varCapabilityActionsMetaData.MaxFwVersion = varCapabilityActionsMetaDataWithoutEmbeddedStruct.MaxFwVersion
		varCapabilityActionsMetaData.MinFwVersion = varCapabilityActionsMetaDataWithoutEmbeddedStruct.MinFwVersion
		varCapabilityActionsMetaData.Model = varCapabilityActionsMetaDataWithoutEmbeddedStruct.Model
		varCapabilityActionsMetaData.SupportedActions = varCapabilityActionsMetaDataWithoutEmbeddedStruct.SupportedActions
		varCapabilityActionsMetaData.TargetType = varCapabilityActionsMetaDataWithoutEmbeddedStruct.TargetType
		varCapabilityActionsMetaData.Vendor = varCapabilityActionsMetaDataWithoutEmbeddedStruct.Vendor
		*o = CapabilityActionsMetaData(varCapabilityActionsMetaData)
	} else {
		return err
	}

	varCapabilityActionsMetaData := _CapabilityActionsMetaData{}

	err = json.Unmarshal(bytes, &varCapabilityActionsMetaData)
	if err == nil {
		o.CapabilityCapability = varCapabilityActionsMetaData.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaxFwVersion")
		delete(additionalProperties, "MinFwVersion")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "SupportedActions")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityActionsMetaData struct {
	value *CapabilityActionsMetaData
	isSet bool
}

func (v NullableCapabilityActionsMetaData) Get() *CapabilityActionsMetaData {
	return v.value
}

func (v *NullableCapabilityActionsMetaData) Set(val *CapabilityActionsMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityActionsMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityActionsMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityActionsMetaData(val *CapabilityActionsMetaData) *NullableCapabilityActionsMetaData {
	return &NullableCapabilityActionsMetaData{value: val, isSet: true}
}

func (v NullableCapabilityActionsMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityActionsMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}