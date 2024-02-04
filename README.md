# golang-scaffold

使用领域领域模型来扩展常见的分层；

方法（包括持久层的逻辑）都定义在model（领域模型）上，之后，这个方法就有很多了。我理解是将这些方法按层次分开；

1. Infrastructure （基础实施层）
    与所有层进行交互 (这里所有层指的是Interfaces，Domain 和 Application )
    譬如：
    1. 自己写的业务工具类
    2. 配置信息
    3. 第三方库的集成和初始化
    4. 数据持久化机制等 
2. Domain （领域层）核心层，业务逻辑会在该层层实现
    包含了
    1. 实体，
    2. 值对象， 
    3. 聚合，
    4. 工厂方法，
    5. Repository仓储实例 
3. Application （应用层）连接domain和 interfaces层；
    对于interface层： 提供各种业务功能方法
    对于domain层： 调用domain层完成任务逻辑；
4. Interfaces （表示层，也叫用户界面层或接口层） 其实就是controller；
负责处理http请求，响应http请求的层；接收参数，校验参数，调业务方法；
可以理解为gin的handler，在脚手架中就是controller


业务层主要调用模型层完成业务的组合调用和事务的封装；
