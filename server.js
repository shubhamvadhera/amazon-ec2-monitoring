var elasticsearch = require('elasticsearch'),
    fs = require('fs'),
    cpu = JSON.parse(fs.readFileSync('CPUUtilization.json'));
    memory = JSON.parse(fs.readFileSync('MemoryUtilization.json'));
    networkIn = JSON.parse(fs.readFileSync('NetworkIn.json'));
    networkOut = JSON.parse(fs.readFileSync('NetworkOut.json'));
    volumeIn = JSON.parse(fs.readFileSync('VolumeReadOps.json'));
    volumeOut = JSON.parse(fs.readFileSync('VolumeWriteOps.json'));
    requestCount = JSON.parse(fs.readFileSync('RequestCount.json'));


var client = new elasticsearch.Client({
  host: 'localhost:9200',
  log: 'trace'
});

if(cpu!=null) {
  for (var a = 0; a < cpu.length; a++ ) {
    client.index({
      index: "cpustats",
      type: "cpuInstanceType",
      id: a,
      body: cpu[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
if(memory != null) {
  for (var a = 0; a < memory.length; a++ ) {
    client.index({
      index: "memorystats",
      type: "memorytype",
      id: a,
      body: memory[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
if(networkIn!= null) {
  for (var a = 0; a < networkIn.length; a++ ) {
    client.index({
      index: "networkinput",
      type: "networkinputtype",
      id: a,
      body: networkIn[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
if(networkOut!=null) {
  for (var a = 0; a < networkOut.length; a++ ) {
    client.index({
      index: "networkoutput",
      type: "networkoutputtype",
      id: a,
      body: networkOut[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
if(volumeIn!=null) {
  for (var a = 0; a < volumeIn.length; a++ ) {
    client.index({
      index: "volumeinput",
      type: "volumeInputType",
      id: a,
      body: volumeIn[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
if(volumeOut!=null) {
  for (var a = 0; a < volumeOut.length; a++ ) {
    client.index({
      index: "volumeoutput",
      type: "volumeOutputType",
      id: a,
      body: volumeOut[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
if(requestCount!=null) {
  for (var a = 0; a < requestCount.length; a++ ) {
    client.index({
      index: "requestcount",
      type: "requestCountType",
      id: a,
      body: requestCount[a]
    }, function(error, response) {
      if (error) {
        console.error(error);
        return;
      }
      else {
      console.log(response);
      }
    });
  }
}
