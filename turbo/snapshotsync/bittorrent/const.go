package bittorrent

import (
	"errors"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/ledgerwatch/turbo-geth/params"
	"github.com/ledgerwatch/turbo-geth/turbo/snapshotsync"
)

const (
	DefaultChunkSize = 1024 * 1024
	SnapshotBlock    = 11_000_000
	LmdbFilename     = "data.mdb"

	HeadersSnapshotHash  = "7727174de470b7fe0bb3e36d35e85cc48853d470" //11кk block 1mb chunk
	BlocksSnapshotHash   = "0546b881c50de9984dd8865d0f18cc5153e4c21b" //11кk block 1mb chunk
	StateSnapshotHash    = ""
	ReceiptsSnapshotHash = ""

	SnapshotInfoHashPrefix  = "ih"
	SnapshotInfoBytesPrefix = "ib"
)

var (
	TorrentHashes = map[uint64]map[snapshotsync.SnapshotType]metainfo.Hash{
		params.MainnetChainConfig.ChainID.Uint64(): {
			snapshotsync.SnapshotType_headers: metainfo.NewHashFromHex(HeadersSnapshotHash),
			snapshotsync.SnapshotType_bodies:  metainfo.NewHashFromHex(BlocksSnapshotHash),
		},
	}
	ErrInvalidSnapshot = errors.New("this snapshot for this chainID not supported ")
)

var Trackers = [][]string{{
	"udp://tracker.openbittorrent.com:80",
	"udp://tracker.openbittorrent.com:80",
	"udp://tracker.publicbt.com:80",
	"udp://coppersurfer.tk:6969/announce",
	"udp://open.demonii.com:1337",
	"http://bttracker.crunchbanglinux.org:6969/announce",
	"udp://wambo.club:1337/announce",
	"udp://tracker.dutchtracking.com:6969/announce",
	"udp://tc.animereactor.ru:8082/announce",
	"udp://tracker.justseed.it:1337/announce",
	"udp://tracker.leechers-paradise.org:6969/announce",
	"udp://tracker.opentrackr.org:1337/announce",
	"https://open.kickasstracker.com:443/announce",
	"udp://tracker.coppersurfer.tk:6969/announce",
	"udp://open.stealth.si:80/announce",
	"http://87.253.152.137/announce",
	"http://91.216.110.47/announce",
	"http://91.217.91.21:3218/announce",
	"http://91.218.230.81:6969/announce",
	"http://93.92.64.5/announce",
	"http://atrack.pow7.com/announce",
	"http://bt.henbt.com:2710/announce",
	"http://bt.pusacg.org:8080/announce",
	"https://tracker.bt-hash.com:443/announce",
	"udp://tracker.leechers-paradise.org:6969",
	"https://182.176.139.129:6969/announce",
	"udp://zephir.monocul.us:6969/announce",
	"https://tracker.dutchtracking.com:80/announce",
	"https://grifon.info:80/announce",
	"udp://tracker.kicks-ass.net:80/announce",
	"udp://p4p.arenabg.com:1337/announce",
	"udp://tracker.aletorrenty.pl:2710/announce",
	"udp://tracker.sktorrent.net:6969/announce",
	"udp://tracker.internetwarriors.net:1337/announce",
	"https://tracker.parrotsec.org:443/announce",
	"https://tracker.moxing.party:6969/announce",
	"https://tracker.ipv6tracker.ru:80/announce",
	"https://tracker.fastdownload.xyz:443/announce",
	"udp://open.stealth.si:80/announce",
	"https://gwp2-v19.rinet.ru:80/announce",
	"https://tr.kxmp.cf:80/announce",
	"https://explodie.org:6969/announce",
}, {
	"udp://zephir.monocul.us:6969/announce",
	"udp://tracker.torrent.eu.org:451/announce",
	"udp://tracker.uw0.xyz:6969/announce",
	"udp://tracker.cyberia.is:6969/announce",
	"http://tracker.files.fm:6969/announce",
	"udp://tracker.zum.bi:6969/announce",
	"http://tracker.nyap2p.com:8080/announce",
	"udp://opentracker.i2p.rocks:6969/announce",
	"udp://tracker.zerobytes.xyz:1337/announce",
	"https://tracker.tamersunion.org:443/announce",
	"https://w.wwwww.wtf:443/announce",
	"https://tracker.imgoingto.icu:443/announce",
	"udp://blokas.io:6969/announce",
	"udp://api.bitumconference.ru:6969/announce",
	"udp://discord.heihachi.pw:6969/announce",
	"udp://cutiegirl.ru:6969/announce",
	"udp://fe.dealclub.de:6969/announce",
	"udp://ln.mtahost.co:6969/announce",
	"udp://vibe.community:6969/announce",
	"http://vpn.flying-datacenter.de:6969/announce",
	"udp://eliastre100.fr:6969/announce",
	"udp://wassermann.online:6969/announce",
	"udp://retracker.local.msn-net.ru:6969/announce",
	"udp://chanchan.uchuu.co.uk:6969/announce",
	"udp://kanal-4.de:6969/announce",
	"udp://handrew.me:6969/announce",
	"udp://mail.realliferpg.de:6969/announce",
	"udp://bubu.mapfactor.com:6969/announce",
	"udp://mts.tvbit.co:6969/announce",
	"udp://6ahddutb1ucc3cp.ru:6969/announce",
	"udp://adminion.n-blade.ru:6969/announce",
	"udp://contra.sf.ca.us:6969/announce",
	"udp://61626c.net:6969/announce",
	"udp://benouworldtrip.fr:6969/announce",
	"udp://sd-161673.dedibox.fr:6969/announce",
	"udp://cdn-1.gamecoast.org:6969/announce",
	"udp://cdn-2.gamecoast.org:6969/announce",
	"udp://daveking.com:6969/announce",
	"udp://bms-hosxp.com:6969/announce",
	"udp://teamspeak.value-wolf.org:6969/announce",
	"udp://edu.uifr.ru:6969/announce",
	"udp://adm.category5.tv:6969/announce",
	"udp://code2chicken.nl:6969/announce",
	"udp://t1.leech.ie:1337/announce",
	"udp://forever-tracker.zooki.xyz:6969/announce",
	"udp://free-tracker.zooki.xyz:6969/announce",
	"udp://public.publictracker.xyz:6969/announce",
	"udp://public-tracker.zooki.xyz:6969/announce",
	"udp://vps2.avc.cx:7171/announce",
	"udp://tracker.fileparadise.in:1337/announce",
	"udp://tracker.skynetcloud.site:6969/announce",
	"udp://z.mercax.com:53/announce",
	"https://publictracker.pp.ua:443/announce",
	"udp://us-tracker.publictracker.xyz:6969/announce",
	"udp://open.stealth.si:80/announce",
	"http://tracker1.itzmx.com:8080/announce",
	"http://vps02.net.orel.ru:80/announce",
	"http://tracker.gbitt.info:80/announce",
	"http://tracker.bt4g.com:2095/announce",
	"https://tracker.nitrix.me:443/announce",
	"udp://aaa.army:8866/announce",
	"udp://tracker.vulnix.sh:6969/announce",
	"udp://engplus.ru:6969/announce",
	"udp://movies.zsw.ca:6969/announce",
	"udp://storage.groupees.com:6969/announce",
	"udp://nagios.tks.sumy.ua:80/announce",
	"udp://tracker.v6speed.org:6969/announce",
	"udp://47.ip-51-68-199.eu:6969/announce",
	"udp://aruacfilmes.com.br:6969/announce",
	"https://trakx.herokuapp.com:443/announce",
	"udp://inferno.demonoid.is:3391/announce",
	"udp://publictracker.xyz:6969/announce",
	"http://tracker2.itzmx.com:6961/announce",
	"http://tracker3.itzmx.com:6961/announce",
	"udp://retracker.akado-ural.ru:80/announce",
	"udp://tracker-udp.gbitt.info:80/announce",
	"http://h4.trakx.nibba.trade:80/announce",
	"udp://tracker.army:6969/announce",
	"http://tracker.anonwebz.xyz:8080/announce",
	"udp://tracker.shkinev.me:6969/announce",
	"http://0205.uptm.ch:6969/announce",
	"udp://tracker.zooki.xyz:6969/announce",
	"udp://forever.publictracker.xyz:6969/announce",
	"udp://tracker.moeking.me:6969/announce",
	"udp://ultra.zt.ua:6969/announce",
	"udp://tracker.publictracker.xyz:6969/announce",
	"udp://ipv4.tracker.harry.lu:80/announce",
	"udp://u.wwwww.wtf:1/announce",
	"udp://line-net.ru:6969/announce",
	"udp://dpiui.reedlan.com:6969/announce",
	"udp://tracker.zemoj.com:6969/announce",
	"udp://t3.leech.ie:1337/announce",
	"http://t.nyaatracker.com:80/announce",
	"udp://exodus.desync.com:6969/announce",
	"udp://valakas.rollo.dnsabr.com:2710/announce",
	"udp://tracker.ds.is:6969/announce",
	"udp://tracker.opentrackr.org:1337/announce",
	"udp://tracker0.ufibox.com:6969/announce",
	"https://tracker.hama3.net:443/announce",
	"udp://opentor.org:2710/announce",
	"udp://t2.leech.ie:1337/announce",
	"https://1337.abcvg.info:443/announce",
	"udp://git.vulnix.sh:6969/announce",
	"udp://retracker.lanta-net.ru:2710/announce",
	"udp://tracker.lelux.fi:6969/announce",
	"udp://bt1.archive.org:6969/announce",
	"udp://admin.videoenpoche.info:6969/announce",
	"udp://drumkitx.com:6969/announce",
	"udp://tracker.dler.org:6969/announce",
	"udp://koli.services:6969/announce",
	"udp://tracker.dyne.org:6969/announce",
	"http://torrenttracker.nwc.acsalaska.net:6969/announce",
	"udp://rutorrent.frontline-mod.com:6969/announce",
	"http://rt.tace.ru:80/announce",
	"udp://explodie.org:6969/announce",
}}
