const express = require('express')
const bodyParser = require('body-parser');
const cors = require('cors');
const app = express()
const k8s = require('@kubernetes/client-node');

const namespace = 'default';
const port = 80


app.use(cors());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());


app.get('/k8ijs/api', (req, res) => {
    console.log(`/api request received`);

    let apiversion = {
        "Name":    "K8ijs",
        "Version": "0.9",
        "Date":    Date().toString(),
    }

    res.json(apiversion);
});


app.get('/k8ijs/list', async (req, res) => {
    console.log(`/list request received`);
    
    const kc = new k8s.KubeConfig();
    kc.loadFromCluster();
    const k8sApi = kc.makeApiClient(k8s.CoreV1Api);

    try {
        console.log(`/list request .... getting info`);
        const podsRes = await k8sApi.listNamespacedPod(namespace);
        const itemsK8s = podsRes.body.items;
        var itemK8sJSON = JSON.stringify(itemsK8s);
        console.log(itemK8sJSON);
        res.json(itemK8sJSON)
    } catch (err) {
        console.error(err);
        res.status(204).send('')
    }

});


app.get('/k8ijs/health', (req, res) => {
    console.log(`/health request received`);
    res.status(200).send('');
});


app.get('/k8ijs/readiness', (req, res) => {
    console.log(`/readiness request received`);
    res.status(200).send('');
});



app.listen(port, () => console.log(`K8ijs rest api running on ${port}!`));