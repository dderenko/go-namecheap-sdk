// Package field is used to log fields with strict types.
// It prevents mapping conflicts in elastic.
package field

import (
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Field = zap.Field

	ObjectMarshaler    = zapcore.ObjectMarshaler
	ObjectEncoder      = zapcore.ObjectEncoder
	ArrayMarshalerFunc = zapcore.ArrayMarshalerFunc
	ArrayEncoder       = zapcore.ArrayEncoder
)

func AbTestId(value uint32) Field                 { return zap.Uint32("ab_test_id", value) }
func Action(actionId uint64) Field                { return zap.String("action", strconv.FormatUint(actionId, 10)) }
func ActionId(actionId uint64) Field              { return zap.Uint64("action_id", actionId) }
func ActionTime(t time.Time) Field                { return zap.String("action_time", t.Format(time.RFC3339)) }
func ActionTimeUnix(t time.Time) Field            { return zap.Int64("action_time_unix", t.Unix()) }
func Address(addr string) Field                   { return zap.String("address", addr) }
func Addresses(addrs []string) Field              { return zap.Strings("addresses", addrs) }
func AdmType(admType fmt.Stringer) Field          { return zap.Stringer("adm_type", admType) }
func AdvertiserId(id int) Field                   { return zap.Int("advertiser_id", id) }
func Args(args []interface{}) Field               { return zap.Reflect("args", args) }
func AudienceId(audienceId uint32) Field          { return zap.Uint32("audience_id", audienceId) }
func BannerId(bannerId int) Field                 { return zap.Int("banner_id", bannerId) }
func BannerRouteImpressionsCount(c int) Field     { return zap.Int("banner_route_impressions_count", c) }
func BannerRouteClicksCount(c int) Field          { return zap.Int("banner_route_clicks_count", c) }
func BannersCount(count int) Field                { return zap.Int("banners_count", count) }
func BidPrice(bitPriceStr string) Field           { return zap.String("bid_price", bitPriceStr) }
func BlackListItemType(blit int) Field            { return zap.Int("black_list_item_type", blit) }
func BranchId(value uint32) Field                 { return zap.Uint32("branch_id", value) }
func BranchSettings(value []byte) Field           { return zap.ByteString("branch_settings", value) }
func BuildDate(date string) Field                 { return zap.String("build date", date) }
func Cache(name string) Field                     { return zap.String("cache", name) }
func CacheKey(key string) Field                   { return zap.String("cache_key", key) }
func CacheName(t string) Field                    { return zap.String("cache_name", t) }
func CacheSize(s int64) Field                     { return zap.Int64("cache_size", s) }
func CampaignActive(b bool) Field                 { return zap.Bool("campaign_active", b) }
func CampaignId(id uint32) Field                  { return zap.Uint32("campaign_id", id) }
func CampaignsCount(count int) Field              { return zap.Int("campaigns_count", count) }
func CheckBy(val string) Field                    { return zap.String("check_by", val) }
func ClicksPath(p string) Field                   { return zap.String("clicks_path", p) }
func ClickType(t string) Field                    { return zap.String("click-type", t) }
func Client(name string) Field                    { return zap.String("client", name) }
func Commit(commit string) Field                  { return zap.String("commit", commit) }
func Component(component string) Field            { return zap.String("component", component) }
func Concurrency(c int) Field                     { return zap.Int("concurrency", c) }
func ConsulService(i interface{}) Field           { return zap.Reflect("consul_service", i) }
func Conversion(conv interface{}) Field           { return zap.Any("conversion", conv) }
func Count(count int) Field                       { return zap.Int("count", count) }
func CountToDelete(count int) Field               { return zap.Int("delete_count", count) }
func CountToUpsert(count int) Field               { return zap.Int("upsert_count", count) }
func CountryCode(countryCode string) Field        { return zap.String("country_code", countryCode) }
func CountryId(countryId uint32) Field            { return zap.Uint32("country_id", countryId) }
func Cpm(value float64) Field                     { return zap.Any("cpm", value) }
func CriteriaId(id uint64) Field                  { return zap.Uint64("criteria_id", id) }
func DebugMCalcMessage(message []byte) Field      { return zap.ByteString("mcalc_message", message) }
func Device(device ObjectMarshaler) Field         { return zap.Object("device", device) }
func DirectionId(actionId uint32) Field           { return zap.Uint32("direction_id", actionId) }
func DirectionPartnerId(id uint32) Field          { return zap.Uint32("direction_partner_id", id) }
func DnsRegistrar(registrar string) Field         { return zap.String("dns_registrar", registrar) }
func Domain(domain string) Field                  { return zap.String("domain", domain) }
func DomainRegistrar(registrar string) Field      { return zap.String("domain_registrar", registrar) }
func DstUrl(url string) Field                     { return zap.String("dst_url", url) }
func DumperIteratorPosition(i int) Field          { return zap.Int("dumper_iterator_position", i) }
func DuplicateCount(c int) Field                  { return zap.Int("duplicate_count", c) }
func Duration(d time.Duration) Field              { return zap.Duration("duration", d) }
func Env(env string) Field                        { return zap.String("env", env) }
func Error(err error) Field                       { return zap.Error(err) }
func ErrorPercent(percent float64) Field          { return zap.Float64("error_percent", percent) }
func Etag(etag string) Field                      { return zap.String("etag", etag) }
func Event(event any) Field                       { return zap.Any("event", event) }
func Exchange(value string) Field                 { return zap.String("exchange", value) }
func ExchangeType(value uint8) Field              { return zap.Uint8("exchange_type", value) }
func ExtendedTestId(id uint64) Field              { return zap.Uint64("extended_test_id", id) }
func FallbackDump(fallbackDump bool) Field        { return zap.Bool("fallback_dump", fallbackDump) }
func Feature(feature string) Field                { return zap.String("feature", feature) }
func FeatureKey(key string) Field                 { return zap.String("feature_key", key) }
func Feed(feed string) Field                      { return zap.String("feed", feed) }
func FeedId(id uint32) Field                      { return zap.Uint32("feed_id", id) }
func FileName(fileName string) Field              { return zap.String("file_name", fileName) }
func GCPercent(percent int) Field                 { return zap.Int("gc_percent", percent) }
func Geo(geo ObjectMarshaler) Field               { return zap.Object("geo", geo) }
func Goal(goal int) Field                         { return zap.Int("goal", goal) }
func GrpcMethod(method string) Field              { return zap.String("grpc.method", method) }
func GrpcPeer(peer string) Field                  { return zap.String("grpc.peer", peer) }
func GrpcService(service string) Field            { return zap.String("grpc.service", service) }
func Handler(handler string) Field                { return zap.String("handler", handler) }
func Height(height int) Field                     { return zap.Int("height", height) }
func Host(host string) Field                      { return zap.String("host", host) }
func HttpBody(body []byte) Field                  { return zap.ByteString("http_body", body) }
func HttpHeaders(headers string) Field            { return zap.String("http_headers", headers) }
func HttpHost(host []byte) Field                  { return zap.ByteString("http_host", host) }
func HttpMethod(method []byte) Field              { return zap.ByteString("http_method", method) }
func HttpRequestUri(uri []byte) Field             { return zap.ByteString("http_request_uri", uri) }
func HttpStatus(status int) Field                 { return zap.Int("http_status", status) }
func Icon(icon string) Field                      { return zap.String("icon", icon) }
func Identity(identity string) Field              { return zap.String("identity", identity) }
func ImpressionsPath(p string) Field              { return zap.String("impressions_path", p) }
func Instance(inst interface{}) Field             { return zap.Reflect("instance", inst) }
func Interval(interval time.Duration) Field       { return zap.Duration("interval", interval) }
func Ip(ip string) Field                          { return zap.String("ip", ip) }
func IsMobile(isMobile bool) Field                { return zap.Bool("is_mobile", isMobile) }
func ItemId(id uint32) Field                      { return zap.Uint32("item_id", id) }
func JsError(kafka string) Field                  { return zap.String("js_error", kafka) }
func Kafka(kafka string) Field                    { return zap.String("kafka", kafka) }
func KafkaName(kafkaName string) Field            { return zap.String("kafka_name", kafkaName) }
func Language(language ObjectMarshaler) Field     { return zap.Object("language", language) }
func Latency(latency uint64) Field                { return zap.Uint64("latency", latency) }
func LenCurrentRoutes(l int) Field                { return zap.Int("len(currentRoutes)", l) }
func LenPreviousRoutes(l int) Field               { return zap.Int("len(previousRoutes)", l) }
func LimiterDiv(val int) Field                    { return zap.Int("limiter_div", val) }
func LimiterZones(val int) Field                  { return zap.Int("limiter_zones", val) }
func LinkedBanners(banners ObjectMarshaler) Field { return zap.Object("linked_banners", banners) }
func Loader(loader string) Field                  { return zap.String("loader", loader) }
func Method(method string) Field                  { return zap.String("method", method) }
func Module(module string) Field                  { return zap.String("module", module) }
func Name(name string) Field                      { return zap.String("name", name) }
func NmsgRequiredBinNumber(n string) Field        { return zap.String("nmsg_required_bin_number", n) }
func NmsgRequiredBinInterval(i []uint32) Field    { return zap.Uint32s("nmsg_required_bin_interval", i) }
func NmsgRequiredBins(bins interface{}) Field     { return zap.Any("nmsg_required_bins", bins) }
func OldValue(value string) Field                 { return zap.String("old_value", value) }
func OsId(osId uint32) Field                      { return zap.Uint32("os_id", osId) }
func Panic(panic interface{}) Field               { return zap.Any("panic", panic) }
func Param(param string) Field                    { return zap.String("param", param) }
func ParamV(name, val string) Field               { return zap.String(fmt.Sprintf("param_%s", name), val) }
func ParamsFromBidRequest(val string) Field       { return zap.String("params_from_bid_request", val) }
func ParamsFromClick(val string) Field            { return zap.String("params_from_click", val) }
func Path(path string) Field                      { return zap.String("path", path) }
func Pid(pid int) Field                           { return zap.Int("pid", pid) }
func Port(port int) Field                         { return zap.Int("port", port) }
func PostbackId(id uint32) Field                  { return zap.Uint32("postback_id", id) }
func PostbacksCount(count int) Field              { return zap.Int("postbacks_count", count) }
func ProcessTime(processTimeMs int64) Field       { return zap.Int64("process_time", processTimeMs) }
func PublisherId(id uint32) Field                 { return zap.Uint32("publisher_id", id) }
func PublisherSiteId(id uint32) Field             { return zap.Uint32("pub_site_id", id) }
func QueryParams(string string) Field             { return zap.String("query_params", string) }
func RateId(rateId uint64) Field                  { return zap.Uint64("rate_id", rateId) }
func RawBytes(data []byte) Field                  { return zap.ByteString("raw_bytes", data) }
func RawData(data []byte) Field                   { return zap.ByteString("raw_data", data) }
func RawExtendedTestCfg(r string) Field           { return zap.String("extended_test_configuration", r) }
func RawMessage(message any) Field                { return zap.Any("raw_message", message) }
func RawPayout(b []byte) Field                    { return zap.ByteString("raw_payout", b) }
func RawPrice(b []byte) Field                     { return zap.ByteString("raw_price", b) }
func RawSettings(raw []byte) Field                { return zap.ByteString("raw_settings", raw) }
func RawStr(data string) Field                    { return zap.String("raw_str", data) }
func RawTimeOnPage(b []byte) Field                { return zap.ByteString("raw_time_on_page", b) }
func RawType(data []byte) Field                   { return zap.ByteString("raw_type", data) }
func Reason(reason string) Field                  { return zap.String("reason", reason) }
func RecencyBinNumber(n string) Field             { return zap.String("recency_bin_number", n) }
func RecencyBinInterval(i []uint32) Field         { return zap.Uint32s("recency_bin_interval", i) }
func RecencyBins(bins interface{}) Field          { return zap.Any("recency_bins", bins) }
func RecordsLimit(l int) Field                    { return zap.Int("records_limit", l) }
func RedirectUrl(url string) Field                { return zap.String("redirect_url", url) }
func Ref(ref string) Field                        { return zap.String("ref", ref) }
func RegressionConfigVersion(v string) Field      { return zap.String("regression_config_version", v) }
func RejectCount(c int) Field                     { return zap.Int("reject_count", c) }
func RejectReason(reason string) Field            { return zap.String("reject_reason", reason) }
func RequestId(id string) Field                   { return zap.String("request_id", id) }
func RequestBody(body []byte) Field               { return zap.ByteString("request_body", body) }
func RequestVar(requestVar string) Field          { return zap.String("request_var", requestVar) }
func Response(resp []byte) Field                  { return zap.String("response", string(resp)) }
func ResponseCache(name string) Field             { return zap.String("response_cache", name) }
func ResultLink(value string) Field               { return zap.String("result_link", value) }
func ResultName(value string) Field               { return zap.String("result_name", value) }
func RotatorRequest(request []byte) Field         { return zap.ByteString("rotator_request", request) }
func RotatorResponse(request []byte) Field        { return zap.ByteString("rotator_response", request) }
func RouteId(id uint32) Field                     { return zap.Uint32("routeId", id) }
func Rps(rps float64) Field                       { return zap.Float64("rps", rps) }
func RpsError(rps float64) Field                  { return zap.Float64("rps_error", rps) }
func Ruid(ruid string) Field                      { return zap.String("ruid", ruid) }
func Seconds(value float64) Field                 { return zap.Float64("seconds", value) }
func SegmentId(segmentId uint32) Field            { return zap.Uint32("segment_id", segmentId) }
func Server(server string) Field                  { return zap.String("server", server) }
func ServerAddressList(list []string) Field       { return zap.Strings("server_address_list", list) }
func Service(service string) Field                { return zap.String("service", service) }
func Sid(value string) Field                      { return zap.String("sid", value) }
func Size(size int) Field                         { return zap.Int("size", size) }
func Source(source string) Field                  { return zap.String("source", source) }
func Stack() Field                                { return zap.Stack("stack") }
func StackTrace(trace []byte) Field               { return zap.ByteString("stack_trace", trace) }
func StatItemKey(key string) Field                { return zap.String("stat_item_key", key) }
func StatusCode(value int) Field                  { return zap.Int("status_code", value) }
func Step(step string) Field                      { return zap.String("step", step) }
func SubZoneId(id uint32) Field                   { return zap.Uint32("sub_zone_id", id) }
func Subsystem(service string) Field              { return zap.String("subsystem", service) }
func TagDomainType(tagDomainType uint32) Field    { return zap.Uint32("tag_domain_type", tagDomainType) }
func Tags(tags []string) Field                    { return zap.Strings("tags", tags) }
func Targeting(targeting []byte) Field            { return zap.ByteString("targeting", targeting) }
func TargetingDictCount(count int) Field          { return zap.Int("targeting_dict_count", count) }
func TargetingDictValuesCount(count int) Field    { return zap.Int("targeting_dict_values_count", count) }
func TargetingKey(value any) Field                { return zap.Any("targeting_key", value) }
func TargetingName(targetingName string) Field    { return zap.String("targeting_name", targetingName) }
func Threshold(threshold uint64) Field            { return zap.Uint64("threshold", threshold) }
func TimeBucket(timeBucket uint8) Field           { return zap.Uint8("time_bucket", timeBucket) }
func Timeout(value time.Duration) Field           { return zap.Duration("timeout", value) }
func To(to string) Field                          { return zap.String("to", to) }
func Token(token string) Field                    { return zap.String("token", token) }
func Topic(topic string) Field                    { return zap.String("topic", topic) }
func Topics(topics any) Field                     { return zap.Any("topics", topics) }
func TplType(t int) Field                         { return zap.Int("tpl_type", t) }
func Trace(value []byte) Field                    { return zap.ByteString("trace", value) }
func TraceId(traceId string) Field                { return zap.String("trace_id", traceId) }
func TrackerId(id int) Field                      { return zap.Int("tracker_id", id) }
func Trx(trx string) Field                        { return zap.String("trx", trx) }
func UniqRoutes(count int) Field                  { return zap.Int("uniq_routes", count) }
func Url(url []byte) Field                        { return zap.ByteString("url", url) }
func UserAgent(ua string) Field                   { return zap.String("user_agent", ua) }
func UserId(userId string) Field                  { return zap.String("user_id", userId) }
func Value(value string) Field                    { return zap.String("value", value) }
func Version(version string) Field                { return zap.String("version", version) }
func Vertical(vertical string) Field              { return zap.String("vertical", vertical) }
func Viewability(v interface{}) Field             { return zap.Any("viewability", v) }
func Worker(worker string) Field                  { return zap.String("worker", worker) }
func Xff(xff string) Field                        { return zap.String("xff", xff) }
func ZoneArg(zoneArg []byte) Field                { return zap.ByteString("zone_arg", zoneArg) }
func ZonesCount(count int) Field                  { return zap.Int("zones_count", count) }
func ZoneId(id uint32) Field                      { return zap.Uint32("zone_id", id) }

var (
	Any     = zap.Any
	String  = zap.String
	Strings = zap.Strings
	Int     = zap.Int
	Ints    = zap.Ints
	Int64   = zap.Int64
	Uint    = zap.Uint
	Uint32  = zap.Uint32
	Uint32s = zap.Uint32s
	Uint64  = zap.Uint64
	Bool    = zap.Bool
	Float64 = zap.Float64
	Time    = zap.Time
	Reflect = zap.Reflect
)
