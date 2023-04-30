// var grpc = require("grpc")

const { isNil } = require("ramda")
// const GRPC_HOST = process.env.GRPC_HOST || "grpc2.weavedb-node.xyz"
// const GRPC_PORT = process.env.GRPC_PORT || "80"
// const GRPC_TLS_PORT = process.env.GRPC_TLS_PORT || "443"
const GRPC_HOST = process.env.GRPC_HOST || "localhost"
const GRPC_PORT = process.env.GRPC_PORT || "3003"

var PROTO_PATH = __dirname + "/weavedb.proto"
var grpc = require("@grpc/grpc-js")
var protoLoader = require("@grpc/proto-loader")
var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
})
var weavedb_proto = grpc.loadPackageDefinition(packageDefinition).weavedb

;(async () => {
  console.log("check nodes")
  const target = GRPC_HOST
  var client = new weavedb_proto.DB(target, grpc.credentials.createInsecure())

  
  // // console.log(client)
  const request = {
    method: "get@o8510D65aT98c3FmiVPgYEzMgSoJDbT6RJwSuTn3NiI",
    query: ["test", 100],
    nocache: true,
  }
  const a = await client.query(request, (err, response) => {
    if (!isNil(err)) {
      console.log("err: ", err)
    }
    return
  })
  console.log(a)

  // client.query(request, function (err, response) {
  //   console.log("query:", response.message)
  // })
  console.log("end")
})()
