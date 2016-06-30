package com.sdu.testprotoreceivewithoutwire;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import proto.hello.HelloOuterClass;

public class MainActivity extends AppCompatActivity {

    /**
     * Server URL (testProtoSend URL)
     */
    public static final String URL = "http://xx.xx.xx.xx:8002/";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        final TextView tx = (TextView) findViewById(R.id.textView);
        final EditText ed = (EditText) findViewById(R.id.editText);

        final Button button = (Button) findViewById(R.id.button);
        if (button != null) {
            button.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    String urlNext = "";
                    if (ed != null) {
                        urlNext = ed.getText().toString();
                    }
                    final Request build = (new Request.Builder()).url(URL + urlNext).build();


                    OkHttpClient okHttpClient = new OkHttpClient.Builder().build();
                    okHttpClient.newCall(build).enqueue(new Callback() {
                        @Override
                        public void onFailure(Call call, IOException e) {

                        }

                        @Override
                        public void onResponse(Call call, Response response) throws IOException {
                            final HelloOuterClass.Hello hello = HelloOuterClass.Hello.parseFrom(response.body().bytes());
                            MainActivity.this.runOnUiThread(new Runnable() {
                                @Override
                                public void run() {
                                    //noinspection ConstantConditions
                                    tx.setText(hello.getName());
                                }
                            });
                        }
                    });

                }

            });

        }
    }
}
