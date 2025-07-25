# TechsupportmanagementTechSupportBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.TechSupportBundle"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.TechSupportBundle"]
**DeviceIdentifier** | Pointer to **string** | The device identifier used to uniquely identify an individual device. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | The device type obtained from the inventory. | [optional] [readonly] 
**Pid** | Pointer to **string** | Product identification of the device. | [optional] 
**PlatformParam** | Pointer to [**NullableConnectorPlatformParamBase**](ConnectorPlatformParamBase.md) |  | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the device. * &#x60;&#x60; - An unrecognized platform type. * &#x60;APIC&#x60; - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * &#x60;CAPIC&#x60; - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * &#x60;DCNM&#x60; - A Cisco Data Center Network Manager (DCNM) instance. * &#x60;UCSFI&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * &#x60;IMC&#x60; - A standalone Cisco UCS rack server (Deprecated). * &#x60;IMCM4&#x60; - A standalone Cisco UCS C-Series or S-Series M4 server. * &#x60;IMCM5&#x60; - A standalone Cisco UCS C-Series or S-Series M5 server. * &#x60;IMCRack&#x60; - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * &#x60;UCSIOM&#x60; - A Cisco UCS Blade Chassis I/O Module (IOM). * &#x60;HX&#x60; - A Cisco HyperFlex (HX) cluster. * &#x60;UCSD&#x60; - A Cisco UCS Director (UCSD) instance. * &#x60;UCSXECMC&#x60; - A Cisco UCSXE Chassis Management Controller. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance instance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist instance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * &#x60;NexusDevice&#x60; - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * &#x60;ACISwitch&#x60; - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * &#x60;NexusSwitch&#x60; - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * &#x60;MDSSwitch&#x60; - A Cisco MDS Switch that is managed using the embedded Device Connector. * &#x60;MDSDevice&#x60; - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * &#x60;UCSC885&#x60; - A standalone Cisco UCS C885 Server. * &#x60;CAI845A&#x60; - A standalone Cisco AI 845A Server. * &#x60;UCSC890&#x60; - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * &#x60;NetAppOntap&#x60; - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVmax&#x60; - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcVplex&#x60; - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;VmwareVcenter&#x60; - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * &#x60;AppDynamics&#x60; - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * &#x60;Dynatrace&#x60; - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;NewRelic&#x60; - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;Umbrella&#x60; - Umbrella cloud target that discovers and monitors an organization. It discovers entities like Datacenters, Devices, Tunnels, Networks, etc. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;MySqlServer&#x60; - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;OracleDatabaseServer&#x60; - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * &#x60;IBMWebSphereApplicationServer&#x60; - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * &#x60;JavaVirtualMachine&#x60; - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * &#x60;AmazonWebService&#x60; - An Amazon Web Service cloud account. Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatform&#x60; - A Google Cloud Platform service account with access to one or more projects. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;GoogleCloudPlatformBilling&#x60; - A Google Cloud Platform service account used to retrieve billing information from BigQuery. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal account with access to Azure subscriptions. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;MicrosoftAzureBilling&#x60; - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * &#x60;DellCompellent&#x60; - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;HPE3Par&#x60; - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * &#x60;NutanixPrismCentral&#x60; - A Nutanix Prism Central cluster. Prism central is a virtual appliance for managing Nutanix clusters and services. * &#x60;HPEOneView&#x60; - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * &#x60;GenericTarget&#x60; - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * &#x60;IMCBlade&#x60; - A Cisco UCS blade server managed by Cisco Intersight. * &#x60;TerraformCloud&#x60; - A Terraform Cloud Business Tier account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * &#x60;CustomTarget&#x60; - CustomTarget is deprecated. Use HTTPEndpoint type to claim HTTP endpoints. * &#x60;AnsibleEndpoint&#x60; - An external endpoint that is added as a target which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * &#x60;HTTPEndpoint&#x60; - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows directly or using Cisco Intersight Assist. Authentication Schemes supported are Basic and Bearer Token. * &#x60;SSHEndpoint&#x60; - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist. * &#x60;CiscoDNAC&#x60; - A Cisco Digital Network Architecture (DNA) Center appliance. * &#x60;CiscoFMC&#x60; - A Cisco Secure Firewall Management Center. * &#x60;ViptelaCloud&#x60; - A Cisco Viptela SD-WAN Cloud. * &#x60;MerakiCloud&#x60; - A Cisco Meraki Organization. * &#x60;CiscoISE&#x60; - A Cisco Identity Services Engine (ISE) target. | [optional] [default to ""]
**Serial** | Pointer to **string** | Serial number of the device. | [optional] 
**DeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**TargetResource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**TechSupportStatus** | Pointer to [**NullableTechsupportmanagementTechSupportStatusRelationship**](TechsupportmanagementTechSupportStatusRelationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementTechSupportBundle

`func NewTechsupportmanagementTechSupportBundle(classId string, objectType string, ) *TechsupportmanagementTechSupportBundle`

NewTechsupportmanagementTechSupportBundle instantiates a new TechsupportmanagementTechSupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportBundleWithDefaults

`func NewTechsupportmanagementTechSupportBundleWithDefaults() *TechsupportmanagementTechSupportBundle`

NewTechsupportmanagementTechSupportBundleWithDefaults instantiates a new TechsupportmanagementTechSupportBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementTechSupportBundle) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementTechSupportBundle) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementTechSupportBundle) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementTechSupportBundle) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementTechSupportBundle) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementTechSupportBundle) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundle) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetDeviceType

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *TechsupportmanagementTechSupportBundle) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetPid

`func (o *TechsupportmanagementTechSupportBundle) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *TechsupportmanagementTechSupportBundle) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *TechsupportmanagementTechSupportBundle) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *TechsupportmanagementTechSupportBundle) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformParam

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformParam() ConnectorPlatformParamBase`

GetPlatformParam returns the PlatformParam field if non-nil, zero value otherwise.

### GetPlatformParamOk

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformParamOk() (*ConnectorPlatformParamBase, bool)`

GetPlatformParamOk returns a tuple with the PlatformParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformParam

`func (o *TechsupportmanagementTechSupportBundle) SetPlatformParam(v ConnectorPlatformParamBase)`

SetPlatformParam sets PlatformParam field to given value.

### HasPlatformParam

`func (o *TechsupportmanagementTechSupportBundle) HasPlatformParam() bool`

HasPlatformParam returns a boolean if a field has been set.

### SetPlatformParamNil

`func (o *TechsupportmanagementTechSupportBundle) SetPlatformParamNil(b bool)`

 SetPlatformParamNil sets the value for PlatformParam to be an explicit nil

### UnsetPlatformParam
`func (o *TechsupportmanagementTechSupportBundle) UnsetPlatformParam()`

UnsetPlatformParam ensures that no value is present for PlatformParam, not even an explicit nil
### GetPlatformType

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *TechsupportmanagementTechSupportBundle) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *TechsupportmanagementTechSupportBundle) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetSerial

`func (o *TechsupportmanagementTechSupportBundle) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *TechsupportmanagementTechSupportBundle) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *TechsupportmanagementTechSupportBundle) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *TechsupportmanagementTechSupportBundle) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundle) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### SetDeviceRegistrationNil

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceRegistrationNil(b bool)`

 SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil

### UnsetDeviceRegistration
`func (o *TechsupportmanagementTechSupportBundle) UnsetDeviceRegistration()`

UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil
### GetTargetResource

`func (o *TechsupportmanagementTechSupportBundle) GetTargetResource() MoBaseMoRelationship`

GetTargetResource returns the TargetResource field if non-nil, zero value otherwise.

### GetTargetResourceOk

`func (o *TechsupportmanagementTechSupportBundle) GetTargetResourceOk() (*MoBaseMoRelationship, bool)`

GetTargetResourceOk returns a tuple with the TargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResource

`func (o *TechsupportmanagementTechSupportBundle) SetTargetResource(v MoBaseMoRelationship)`

SetTargetResource sets TargetResource field to given value.

### HasTargetResource

`func (o *TechsupportmanagementTechSupportBundle) HasTargetResource() bool`

HasTargetResource returns a boolean if a field has been set.

### SetTargetResourceNil

`func (o *TechsupportmanagementTechSupportBundle) SetTargetResourceNil(b bool)`

 SetTargetResourceNil sets the value for TargetResource to be an explicit nil

### UnsetTargetResource
`func (o *TechsupportmanagementTechSupportBundle) UnsetTargetResource()`

UnsetTargetResource ensures that no value is present for TargetResource, not even an explicit nil
### GetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundle) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship`

GetTechSupportStatus returns the TechSupportStatus field if non-nil, zero value otherwise.

### GetTechSupportStatusOk

`func (o *TechsupportmanagementTechSupportBundle) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool)`

GetTechSupportStatusOk returns a tuple with the TechSupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundle) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship)`

SetTechSupportStatus sets TechSupportStatus field to given value.

### HasTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundle) HasTechSupportStatus() bool`

HasTechSupportStatus returns a boolean if a field has been set.

### SetTechSupportStatusNil

`func (o *TechsupportmanagementTechSupportBundle) SetTechSupportStatusNil(b bool)`

 SetTechSupportStatusNil sets the value for TechSupportStatus to be an explicit nil

### UnsetTechSupportStatus
`func (o *TechsupportmanagementTechSupportBundle) UnsetTechSupportStatus()`

UnsetTechSupportStatus ensures that no value is present for TechSupportStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


