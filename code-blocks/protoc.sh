protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/person.proto
# Because you want Go classes, you use the --go_out option –