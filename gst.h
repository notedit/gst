#include <stdlib.h>
#include <gst/gst.h>
#include <gst/app/gstappsrc.h>
#include <gst/app/gstappsink.h>
#include <string.h>

extern void X_gst_shim_init();
extern void X_gst_bin_add(GstElement* p, GstElement* element1);
extern void X_gst_bin_remove(GstElement *p, GstElement *element);
extern void X_gst_g_object_set_string(GstElement *e, const gchar* p_name, gchar* p_value);
extern void X_gst_g_object_set_int(GstElement *e, const gchar* p_name, gint p_value);
extern void X_gst_g_object_set_uint(GstElement *e, const gchar* p_name, guint p_value);
extern void X_gst_g_object_set_bool(GstElement *e, const gchar* p_name, gboolean p_value);
extern void X_gst_g_object_set_caps(GstElement *e, const gchar* p_name, const GstCaps *p_value);
extern void X_gst_g_object_set(GstElement* e, const gchar* p_name, const GValue* p_value);
extern void X_gst_g_object_set_structure(GstElement *e, const gchar* p_name, const GstStructure *p_value);
extern void X_gst_g_object_setv(GObject* object, guint n_properties, const gchar* names[], const GValue value[]);
extern void X_g_signal_connect(GstElement* element, gchar* detailed_signal, void (*f)(GstElement*, GstPad*, gpointer), gpointer data);
extern GstElement* cb_request_aux_receiver(GstElement *rtpbin, guint session, gpointer user_data);
extern GstElement* cb_request_aux_sender(GstElement *rtpbin, guint session, gpointer user_data);
extern void cb_new_pad(GstElement* element, GstPad* pad, gpointer data);
extern void cb_need_data(GstElement *element, guint size, gpointer data);
extern void cb_enough_data(GstElement *element, gpointer data);
extern gboolean cb_pad_event(GstPad *pad, GstObject *parent, GstEvent *event);
extern GstElement *X_gst_bin_get_by_name(GstElement* element, const gchar* name);
extern GstElementClass *X_GST_ELEMENT_GET_CLASS(GstElement *element);
extern void X_GST_DEBUG_BIN_TO_DOT_FILE(GstElement *element, const gchar* name);
extern void X_g_signal_emit_buffer_by_name(GstElement* element, const gchar* detailed_signal, GstBuffer* buffer, GstFlowReturn* ret);
extern GstBuffer *X_gst_buffer_new_wrapped(gchar* src, gsize len);
extern gboolean X_gst_buffer_map(GstBuffer* gstBuffer, GstMapInfo* mapInfo);
extern void X_gst_pipeline_use_clock(GstElement *element, GstClock *clock);
extern void X_gst_element_set_start_time_none(GstElement *element);
extern void X_gst_structure_set_string(GstStructure *structure, const gchar *name, gchar* value);
extern void X_gst_structure_set_int(GstStructure *structure, const gchar *name, gint value);
extern void X_gst_structure_set_uint(GstStructure *structure, const gchar *name, guint value);
extern void X_gst_structure_set_bool(GstStructure *structure, const gchar *name, gboolean value);
extern GstEventType X_GST_EVENT_TYPE(GstEvent* event);
extern GstMessageType X_GST_MESSAGE_TYPE(GstMessage *message);
extern GstBus* X_gst_pipeline_get_bus(GstElement* element);
extern void cb_bus_message(GstBus * bus, GstMessage * message, gpointer poll_data);
extern void X_g_signal_connect_data(gpointer instance, const gchar *detailed_signal, void (*f)(GstElement*, GstBus*, GstMessage*, gpointer), gpointer data, GClosureNotify destroy_data, GConnectFlags connect_flags);
