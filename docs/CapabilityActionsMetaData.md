# CapabilityActionsMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ActionsMetaData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ActionsMetaData"]
**MaxFwVersion** | Pointer to **string** | Maximum firmware version of the server. | [optional] 
**MinFwVersion** | Pointer to **string** | Minimum firmware version of the server. | [optional] 
**Model** | Pointer to **string** | Model of the server that supports power or storage operations. | [optional] 
**SupportedActions** | Pointer to **[]string** |  | [optional] 
**TargetType** | Pointer to **string** | The target type to which the metadata applies. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;CAPIC&#x60; - A Cloud Application Policy Infrastructure Controller instance. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;IMCRack&#x60; - A standalone UCS M6 and above server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;IWE&#x60; - An Intersight Workload Engine. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NexusDevice&#x60; - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * &#x60;ACISwitch&#x60; - A platform type to support ACI Switches. * &#x60;NexusSwitch&#x60; - A platform type to support Cisco Nexus Switches. * &#x60;MDSSwitch&#x60; - A platform type to support Cisco MDS Switches. * &#x60;MDSDevice&#x60; - A platform type to support MDS devices. * &#x60;UCSC890&#x60; - A standalone Cisco UCSC890 server. * &#x60;RedfishServer&#x60; - A generic target type for servers that support Redfish. Current support is limited to managing HPE and Dell servers on Intersight. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;NewRelic&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;ReadHatOpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;MySqlServer&#x60; - An instance of either Oracle MySQL Database or the open source MariaDB. * &#x60;OracleDatabaseServer&#x60; - The Oracle Server is a relational database management system that provides an open, comprehensive, and integrated approach to information management. * &#x60;IBMWebSphereApplicationServer&#x60; - WebSphere Application Server (WAS) is a software product that performs the role of a web application server. More specifically it is a software framework and middleware that hosts Java based web applications. * &#x60;OracleWebLogicServer&#x60; - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * &#x60;ApacheTomcatServer&#x60; - Apache Tomcat is a web container. It allows the users to run Servlet and JAVA Server Pages that are based on the web-applications. * &#x60;JavaVirtualMachine&#x60; - The Java Virtual Machine (JVM) is the runtime engine of the Java Platform, which allows any program written in Java or other language compiled into Java bytecode to run on any computer that has a native JVM. * &#x60;RedHatJBossApplicationServer&#x60; - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;AmazonWebService&#x60; - An Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * &#x60;AmazonWebServiceBilling&#x60; - An Amazon web service billing target to retrieve billing information stored in S3 bucket. * &#x60;GoogleCloudPlatform&#x60; - Google Cloud Platform (GCP), offered by Google, is a suite of cloud computing services that runs on the same infrastructure that Google uses internally for its end-user products, such as Google Search, Gmail, Google Drive, and YouTube. Alongside a set of management tools, it provides a series of modular cloud services including computing, data storage, data analytics and machine learning. Google Cloud Platform provides infrastructure as a service, platform as a service, and serverless computing environments. * &#x60;GoogleCloudPlatformBilling&#x60; - Google Cloud Platform (GCP) offers flexible ways to set up and manage billing for your resources. A billing account is how a user pays for the resources being consumed. A billing account is associated with a method of payment and access is established using Cloud IAM roles. For a resource to be deployed in a project, the project has to be associated with a billing account. More than one project can be associated with a billing account. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * &#x60;DellCompellent&#x60; - A Dell Compellent storage system. * &#x60;HPE3Par&#x60; - A HPE 3PAR storage system. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * &#x60;HPEOneView&#x60; - A HPE Oneview management system that manages compute, storage, and networking. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. * &#x60;TerraformCloud&#x60; - A Terraform Cloud account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * &#x60;CustomTarget&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;AnsibleEndpoint&#x60; - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * &#x60;HTTPEndpoint&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * &#x60;SSHEndpoint&#x60; - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows machine on which PowerShell scripts can be executed remotely. | [optional] [default to ""]
**Vendor** | Pointer to **string** | Manufacturer of the server. | [optional] 

## Methods

### NewCapabilityActionsMetaData

`func NewCapabilityActionsMetaData(classId string, objectType string, ) *CapabilityActionsMetaData`

NewCapabilityActionsMetaData instantiates a new CapabilityActionsMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityActionsMetaDataWithDefaults

`func NewCapabilityActionsMetaDataWithDefaults() *CapabilityActionsMetaData`

NewCapabilityActionsMetaDataWithDefaults instantiates a new CapabilityActionsMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityActionsMetaData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityActionsMetaData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityActionsMetaData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityActionsMetaData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityActionsMetaData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityActionsMetaData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaxFwVersion

`func (o *CapabilityActionsMetaData) GetMaxFwVersion() string`

GetMaxFwVersion returns the MaxFwVersion field if non-nil, zero value otherwise.

### GetMaxFwVersionOk

`func (o *CapabilityActionsMetaData) GetMaxFwVersionOk() (*string, bool)`

GetMaxFwVersionOk returns a tuple with the MaxFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFwVersion

`func (o *CapabilityActionsMetaData) SetMaxFwVersion(v string)`

SetMaxFwVersion sets MaxFwVersion field to given value.

### HasMaxFwVersion

`func (o *CapabilityActionsMetaData) HasMaxFwVersion() bool`

HasMaxFwVersion returns a boolean if a field has been set.

### GetMinFwVersion

`func (o *CapabilityActionsMetaData) GetMinFwVersion() string`

GetMinFwVersion returns the MinFwVersion field if non-nil, zero value otherwise.

### GetMinFwVersionOk

`func (o *CapabilityActionsMetaData) GetMinFwVersionOk() (*string, bool)`

GetMinFwVersionOk returns a tuple with the MinFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFwVersion

`func (o *CapabilityActionsMetaData) SetMinFwVersion(v string)`

SetMinFwVersion sets MinFwVersion field to given value.

### HasMinFwVersion

`func (o *CapabilityActionsMetaData) HasMinFwVersion() bool`

HasMinFwVersion returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityActionsMetaData) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityActionsMetaData) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityActionsMetaData) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityActionsMetaData) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSupportedActions

`func (o *CapabilityActionsMetaData) GetSupportedActions() []string`

GetSupportedActions returns the SupportedActions field if non-nil, zero value otherwise.

### GetSupportedActionsOk

`func (o *CapabilityActionsMetaData) GetSupportedActionsOk() (*[]string, bool)`

GetSupportedActionsOk returns a tuple with the SupportedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedActions

`func (o *CapabilityActionsMetaData) SetSupportedActions(v []string)`

SetSupportedActions sets SupportedActions field to given value.

### HasSupportedActions

`func (o *CapabilityActionsMetaData) HasSupportedActions() bool`

HasSupportedActions returns a boolean if a field has been set.

### SetSupportedActionsNil

`func (o *CapabilityActionsMetaData) SetSupportedActionsNil(b bool)`

 SetSupportedActionsNil sets the value for SupportedActions to be an explicit nil

### UnsetSupportedActions
`func (o *CapabilityActionsMetaData) UnsetSupportedActions()`

UnsetSupportedActions ensures that no value is present for SupportedActions, not even an explicit nil
### GetTargetType

`func (o *CapabilityActionsMetaData) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CapabilityActionsMetaData) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CapabilityActionsMetaData) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CapabilityActionsMetaData) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetVendor

`func (o *CapabilityActionsMetaData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityActionsMetaData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityActionsMetaData) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityActionsMetaData) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

