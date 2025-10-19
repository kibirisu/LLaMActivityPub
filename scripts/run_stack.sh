#!/usr/bin/env bash

LOG_DIR=/tmp/
start_backend() {
    if [ -n "$BACKEND_PID" ]; then
        kill $BACKEND_PID 2>/dev/null
    fi
    pushd server > /dev/null
    go build main.go
    ./main > $BACKEND_LOGS &
    BACKEND_PID=$!
    echo "Backend started with pid: $BACKEND_PID"
    popd > /dev/null
    trap "pg_ctl -D \"$PGDATA\" stop || true;
        kill $BACKEND_PID 2>/dev/null || true;
        kill $FRONTEND_PID 2>/dev/null || true;" EXIT
}

start_frontend() {
    if [ -n "$FRONTEND_PID" ]; then
        kill $FRONTEND_PID 2>/dev/null
    fi
    pushd web  > /dev/null
    BROWSER=none pnpm run dev > $FRONTEND_LOGS &
    FRONTEND_PID=$!
    echo "Frontend started with pid: $FRONTEND_PID"
    popd > /dev/null
    trap "pg_ctl -D \"$PGDATA\" stop || true;
        kill $BACKEND_PID 2>/dev/null || true;
        kill $FRONTEND_PID 2>/dev/null || true;" EXIT
}

start_postgres() {
    PGDATA=$PWD/.pgdata
    PGSOCKET_DIR="$PWD/.pgsocket"
    mkdir -p $PGSOCKET_DIR
    PGPORT=5432
    if [ ! -d "$PGDATA" ]; then
      initdb -D "$PGDATA"
    fi
    pg_ctl -D "$PGDATA" -o "-p $PGPORT -k $PGSOCKET_DIR" -l $LOG_DIR/postgresql.log start
    echo "PostgreSQL started on port $PGPORT"
    trap "pg_ctl -D \"$PGDATA\" stop || true;
        kill $BACKEND_PID 2>/dev/null || true;
        kill $FRONTEND_PID 2>/dev/null || true;" EXIT
}

if [ -z "$BACKEND_LOGS" ]; then
    BACKEND_LOGS="$LOG_DIR/backend.log"
fi
FRONTEND_LOGS=$2
if [ -z "$FRONTEND_LOGS" ]; then
    FRONTEND_LOGS="$LOG_DIR/frontend.log"
fi

start_postgres
start_backend
start_frontend


echo "Press 'f' to restart frontend, 'b' to restart backend, or 'q' to quit"
while true; do
    # Read one key silently
    read -n 1 -s key

    case "$key" in
        f)
            echo "Restarting frontend..."
            start_frontend
            ;;
        b)
            echo "Restarting backend..."
            start_backend
            ;;
        q)
            echo "Exiting."
            break
            ;;
        *)
            echo "Unknown key: $key (use 'q', 'b' or 'f')"
            ;;
    esac
done
echo

