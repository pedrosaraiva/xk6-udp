import udp from 'k6/x/udp';
import { check } from 'k6';

const conn = udp.connect('host:port');

export default function () {
    udp.writeLn(conn, 'Say Hello');
    let res = String.fromCharCode(...udp.read(conn, 1024))
    check(res, {
        'verify ag tag': (res) => res.includes('Hello')
    });
    udp.close(conn);
}