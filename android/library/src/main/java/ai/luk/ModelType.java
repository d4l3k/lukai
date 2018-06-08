package ai.luk;

import java.util.Map;
import java.util.List;
import java.io.ByteArrayOutputStream;
import java.io.UnsupportedEncodingException;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;
import java.lang.String;

public class ModelType {
  public ai.luk.bindable.ModelType mt;

  public ModelType(String domain, String modelType, String dataDir) throws Exception {
    mt = ai.luk.bindable.Bindable.makeModelType(domain, modelType, dataDir);
  }

  public boolean isTraining() {
    return mt.isTraining();
  }

  public void startTraining() throws Exception {
    mt.startTraining();
  }

  public void stopTraining() throws Exception {
    mt.stopTraining();
  }

  public long totalExamples() {
    return mt.totalExamples();
  }

  public void log(Map<String, Tensor> feeds, List<String> targets) throws Exception {
    mt.log(serialize(feeds), serialize(targets));
  }

  public Map<String, Tensor> run(
      Map<String, Tensor> feeds, List<String> fetches, List<String> targets) throws Exception {
    return unserializeMap(mt.run(serialize(feeds), serialize(fetches), serialize(targets)));
  }

  public static void serializeTo(ByteArrayOutputStream w, long n) {
    ByteBuffer buffer = ByteBuffer.allocate(Long.SIZE / Byte.SIZE).order(ByteOrder.nativeOrder());
    buffer.putLong(n);
    byte[] body = buffer.array();
    w.write(body, 0, body.length);
  };

  public static void serializeTo(ByteArrayOutputStream w, float n) {
    ByteBuffer buffer = ByteBuffer.allocate(Float.SIZE / Byte.SIZE).order(ByteOrder.nativeOrder());
    buffer.putFloat(n);
    byte[] body = buffer.array();
    w.write(body, 0, body.length);
  };

  public static void serializeTo(ByteArrayOutputStream w, double n) {
    ByteBuffer buffer = ByteBuffer.allocate(Double.SIZE / Byte.SIZE).order(ByteOrder.nativeOrder());
    buffer.putDouble(n);
    byte[] body = buffer.array();
    w.write(body, 0, body.length);
  };

  public static void serializeTo(ByteArrayOutputStream w, int n) {
    ByteBuffer buffer = ByteBuffer.allocate(Integer.SIZE / Byte.SIZE).order(ByteOrder.nativeOrder());
    buffer.putInt(n);
    byte[] body = buffer.array();
    w.write(body, 0, body.length);
  };

  public static void serializeTo(ByteArrayOutputStream w, Integer n) {
    serializeTo(w, (int)n);
  };

  public static void serializeTo(ByteArrayOutputStream w, Float n) {
    serializeTo(w, (float)n);
  };

  public static void serializeTo(ByteArrayOutputStream w, Long n) {
    serializeTo(w, (long)n);
  };

  public static void serializeTo(ByteArrayOutputStream w, Double n) {
    serializeTo(w, (double)n);
  };

  public static void serializeTo(ByteArrayOutputStream w, long[] arr) {
    serializeTo(w, (long) arr.length);
    for (long item : arr) {
      serializeTo(w, item);
    }
  };

  public static void serializeTo(ByteArrayOutputStream w, String str) {
    try {
      byte[] body = str.getBytes("UTF-8");
      serializeTo(w, (long) body.length);
      w.write(body, 0, body.length);
    } catch (UnsupportedEncodingException e) {
      throw new RuntimeException(e);
    }
  }

  public static void serializeTo(ByteArrayOutputStream w, Tensor t) {
    serializeTo(w, (long) t.dataType().c());
    serializeTo(w, t.shape());
    byte[] body = t.bytesValue();
    serializeTo(w, (long) body.length);
    w.write(body, 0, body.length);
  }

  public static <T> void serializeTo(ByteArrayOutputStream w, List<T> list) {
    serializeTo(w, (long) list.size());
    for (T item : list) {
      serializeTo(w, item);
    }
  }

  public static void serializeTo(ByteArrayOutputStream w, Object unknown) {
    if (unknown instanceof String) {
      serializeTo(w, (String)unknown);
    } else if (unknown instanceof List) {
      serializeTo(w, (List)unknown);
    } else if (unknown instanceof Integer) {
      serializeTo(w, (Integer)unknown);
    } else if (unknown instanceof Long) {
      serializeTo(w, (Long)unknown);
    } else if (unknown instanceof Float) {
      serializeTo(w, (Float)unknown);
    } else if (unknown instanceof Double) {
      serializeTo(w, (Double)unknown);
    } else {
      throw new RuntimeException("unknown type serialize: " + unknown.getClass().getName());
    }
  }

  public static byte[] serialize(List<?> list) {
    ByteArrayOutputStream w = new ByteArrayOutputStream();
    serializeTo(w, list);
    return w.toByteArray();
  }

  public static byte[] serialize(Map<String, Tensor> list) throws UnsupportedEncodingException {
    ByteArrayOutputStream w = new ByteArrayOutputStream();
    serializeTo(w, (long) list.size());
    for (Map.Entry<String, Tensor> entry : list.entrySet()) {
      serializeTo(w, entry.getKey());
      serializeTo(w, entry.getValue());
    }
    return w.toByteArray();
  }

  public static Map<String, Tensor> unserializeMap(byte[] body) {
    return null;
  }
}
