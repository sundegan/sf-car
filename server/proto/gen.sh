PBTS_OUT_DIR=../../wx/miniprogram/service/proto_gen

pbjs -t static -w es6 trip.proto \
--no-create --no-encode --no-decode --no-verify --no-delimited --force-number \
-o $PBTS_OUT_DIR/trip_pb_tmp.js 
echo 'import * as $protobuf from "protobufjs";\n' > $PBTS_OUT_DIR/trip_pb.js
cat $PBTS_OUT_DIR/trip_pb_tmp.js >> $PBTS_OUT_DIR/trip_pb.js
rm $PBTS_OUT_DIR/trip_pb_tmp.js
pbts -o $PBTS_OUT_DIR/trip_pb.d.ts $PBTS_OUT_DIR/trip_pb.js