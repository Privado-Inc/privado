# Configurations

Privado’s execution can be configured in multiple ways. All configurations are located under the config folder. 

# Metrics Collection

Privado collects performance metrics about its own execution. These metrics do not identify you and are completely anonymous. If you still have concerns, you can turn this off by opening the file `$HOME/.privado.config.json` and setting the following values:

    {
        "metrics": false,
        "syncToPrivadoCloud": true
    }      
    
   
# Privado Cloud Sync

By default, Privado results are synced to the Privado Cloud so that users can use the Privado Dashboard for reviewing results, observing data flows and analyzing privacy issues. **No code is sent to the cloud.** Only the results summary is sent for visualization. If you do not wish to use Privado cloud, you can turn this off by opening the file `$HOME/.privado.config.json` and setting the following values:

    {
        "metrics": true,
        "syncToPrivadoCloud": false
    }    




