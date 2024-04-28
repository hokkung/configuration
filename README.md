# Configuration

# Scope
1. Able to update configuration through APIs
2. This service will provide APIs to retreiving configuration
3. This service will store data into Redis as database
4. Destinations must implement cache before sending request to this service
5. After updating the configuration the service will be sending updated value to subscribed services through APIs/Message queue


# Entity
1. Service
2. Configuration


# Internal Libs
1. [srv](https://github.com/hokkung/srv)